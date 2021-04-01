/*
 * Cloud API
 *
 * The public facing API through which connectors are exposed as a single abtract API
 *
 * API version: v1.5
 * Contact: support@trexis.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finite

// FiniteEvent : Type of events supported by Finite
type FiniteEvent string

// List of FiniteEvent
const (
	CACHE_DELETE FiniteEvent = "CacheDelete"
	CACHE_UPDATE FiniteEvent = "CacheUpdate"
	CACHE_READ   FiniteEvent = "CacheRead"
	ITEM_UPDATE  FiniteEvent = "ItemUpdate"
	ITEM_READ    FiniteEvent = "ItemRead"
	ITEM_DELETE  FiniteEvent = "ItemDelete"
	LIST_READ    FiniteEvent = "ListRead"
)
