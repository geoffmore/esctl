package esutil

// This should be used as a basis to modify json objects to expose higher level functions like adding aggregations to existing queries
// What types are exposed in json? How do they map to what's available in Golang?
/*
 JSON: 		golang
 string: 	string
 object: map[string]interface{}?
 array: list
 boolean: boolean
 null: nil

 */

// Function ideas
// merge(obj1, obj2) - combine two objects, de-duplicate keys
// insert(obj2, obj1, key) insert an object into another object at the specific key

// Where can I find the JSON operations used by kubectl?
// https://github.com/kubernetes/kubectl/tree/master/pkg/apply/strategy
// https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#patch
// Three types of patches are:
/*
	StrategicMergePatch (https://github.com/kubernetes/community/blob/master/contributors/devel/sig-api-machinery/strategic-merge-patch.md)
	JSONMergePatch (https://tools.ietf.org/html/rfc7386)
	JSONPatch (https://tools.ietf.org/html/rfc6902)
 */