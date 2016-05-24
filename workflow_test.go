package workflow

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseInfo(t *testing.T) {
	info := DefaultWorkflow().Info()
	if info.BundleID != "net.deanishe.awgo" {
		t.Fatalf("Incorrect bundle ID: %v", info.BundleID)
	}

	if info.Author != "Dean Jackson" {
		t.Fatalf("Incorrect author: %v", info.Author)
	}

	if info.Description != "awgo sample info.plist" {
		t.Fatalf("Incorrect description: %v", info.Description)
	}

	if info.Name != "awgo" {
		t.Fatalf("Incorrect name: %v", info.Name)
	}

	if info.Website != "https://gogs.deanishe.net/deanishe/awgo" {
		t.Fatalf("Incorrect website: %v", info.Website)
	}
}

// TestParseVars tests that variables are read from info.plist
func TestParseVars(t *testing.T) {
	i := DefaultWorkflow().Info()
	if i.Var("exported_var") != "exported_value" {
		t.Fatalf("exported_var=%v, expected=exported_value", i.Var("exported_var"))
	}

	// Should unexported variables be ignored?
	if i.Var("unexported_var") != "unexported_value" {
		t.Fatalf("unexported_var=%v, expected=unexported_value", i.Var("unexported_var"))
	}
}

func ExampleInfo_Var() {
	i := GetInfo()
	fmt.Println(i.Var("exported_var"))
	// Output: exported_value
}

func ExampleNewWorkflow() {
	wf := NewWorkflow(&Options{})
	// Version is read from info.plist
	fmt.Println(wf.Version())
	// Output: 0.2.1
}

func ExampleNewWorkflow_overrideVersion() {
	// Override the version string read from info.plist (if present)
	wf := NewWorkflow(&Options{Version: "1.1.0"})
	fmt.Println(wf.Version())
	// Output: 1.1.0
}

// The normal way to create a new Item, but not the normal way to use it.
//
// Normally, when you're done adding Items, you call SendFeedback() to
// send the results to Alfred.
func ExampleNewItem() {
	// Create a new item via the default Workflow object, which will
	// track the Item and send it to Alfred when you call SendFeedback()
	it := NewItem()
	// Title is required
	it.Title = "First Result"
	it.Subtitle = "Some details here"
	// Don't know why you'd do this, but...
	data, _ := json.Marshal(it)
	fmt.Println(string(data))
	// Output: {"title":"First Result","subtitle":"Some details here"}
}
