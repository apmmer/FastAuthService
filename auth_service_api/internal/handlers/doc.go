// This package is responsible for requests and responses handling.
//
// Each file in this package contains a function related to a specific endpoint.
// If some functionality can be the same for different handlers,
// then such functionality is carried out to the handlers_utils package.

// Usually handlers use 'private repositories' to create or
// retrieve serialized objects. All the functionality of the handlers
// is aimed at directly processing the data received with the request,
// possibly validating, converting and preparing this data,
// for further sending to the 'repository' to perform operations
// on the prepared data.
package handlers
