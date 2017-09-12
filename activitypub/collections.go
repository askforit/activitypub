package activitypub

type Collection struct {
	*BaseObject
	// A non-negative integer specifying the total number of objects contained by the logical view of the collection.
	// This number might not reflect the actual number of items serialized within the Collection object instance.
	TotalItems uint				`jsonld:"totalItems,omitempty"`
	// Identifies the items contained in a collection. The items might be ordered or unordered.
	Items      ItemCollection	`jsonld:"items,omitempty"`
}

type OrderedCollection struct {
	*BaseObject
	// A non-negative integer specifying the total number of objects contained by the logical view of the collection.
	// This number might not reflect the actual number of items serialized within the Collection object instance.
	TotalItems uint				`jsonld:"totalItems,omitempty"`
	// Identifies the items contained in a collection. The items might be ordered or unordered.
	Items      ItemCollection	`jsonld:"items,omitempty"`
}

type Page ObjectOrLink

type CollectionPage struct {
	PartOf *Collection
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current Page `jsonld:"current,omitempty"`
	// In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First   Page `jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last    Page `jsonld:"last,omitempty"`
	// In a paged Collection, indicates the next page of items.
	Next    Page `jsonld:"next,omitempty"`
	// In a paged Collection, identifies the previous page of items.
	Prev    Page `jsonld:"prev,omitempty"`
}

type OrderedCollectionPage struct {
	PartOf *OrderedCollection
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current 	Page 	`jsonld:"current,omitempty"`
	// In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First   	Page 	`jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last    	Page 	`jsonld:"last,omitempty"`
	// In a paged Collection, indicates the next page of items.
	Next    	Page 	`jsonld:"next,omitempty"`
	// In a paged Collection, identifies the previous page of items.
	Prev    	Page 	`jsonld:"prev,omitempty"`
	// A non-negative integer value identifying the relative position within the logical view of a strictly ordered collection.
	StartIndex	uint	`jsonld:"startIndex,omitempty"`
}


func CollectionNew(id ObjectId) *Collection {
	o := BaseObject{Id: id, Type: CollectionType}

	return &Collection{BaseObject:&o}
}

func OrderedCollectionNew(id ObjectId) *OrderedCollection {
	o := BaseObject{Id: id, Type:  OrderedCollectionType}

	return &OrderedCollection{BaseObject:&o}
}

func CollectionPageNew(parent *Collection) *CollectionPage {
	return &CollectionPage{PartOf: parent}
}
func OrderedCollectionPageNew(parent *OrderedCollection) *OrderedCollectionPage {
	return &OrderedCollectionPage{PartOf: parent}
}