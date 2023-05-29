// This package contains utils (functions) that can be used only by repositories.

// All repositories except the base one (base_repo) work with data transformation
// and a specific model related to this repository.
// All database requests are made using the base repository.
// Many auxiliary functions that will be identical and may occur in
// different repositories will be located in this module.
package repositories_utils
