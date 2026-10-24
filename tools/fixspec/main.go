// fixspec downloads an OpenAPI spec and prepares it for code generation.
//
// The Bluestone specs mark most optional fields as `nullable`. oapi-codegen
// generates those without `omitempty`, so every unset field is sent as an
// explicit `null`. For PATCH-like endpoints that changes the meaning of the
// request (`"name": null` clears the name instead of leaving it untouched).
//
// To keep "unset means omitted", `nullable` is removed from the spec. The only
// exception is the `value` of the PropertyUpdate* wrapper objects: there an
// explicit `{"value": null}` is how a field is cleared, so it is kept nullable
// (and thus generated without `omitempty`).
//
// Usage: go run ./tools/fixspec <url> <output file> (see `task download-specs`)
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: fixspec <url> <output file>")
		os.Exit(2)
	}

	if err := run(os.Args[1], os.Args[2]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(url, output string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetching %s: unexpected status %d", url, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var spec map[string]any
	if err := json.Unmarshal(body, &spec); err != nil {
		return fmt.Errorf("parsing %s: %w", url, err)
	}

	removeNullable(spec)
	markPropertyUpdateValuesNullable(spec)

	result, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(output, append(result, '\n'), 0o644)
}

func removeNullable(v any) {
	switch node := v.(type) {
	case map[string]any:
		delete(node, "nullable")
		for _, child := range node {
			removeNullable(child)
		}
	case []any:
		for _, child := range node {
			removeNullable(child)
		}
	}
}

func markPropertyUpdateValuesNullable(spec map[string]any) {
	components, _ := spec["components"].(map[string]any)
	schemas, _ := components["schemas"].(map[string]any)

	for name, schema := range schemas {
		if !strings.HasPrefix(name, "PropertyUpdate") {
			continue
		}
		s, _ := schema.(map[string]any)
		properties, _ := s["properties"].(map[string]any)
		if value, ok := properties["value"].(map[string]any); ok {
			value["nullable"] = true
		}
	}
}
