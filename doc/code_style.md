# Code Style

* A comment with the filename is the first line in each file
* The short version of the license is at the top of each file
* Code is sorted in each file by keyword, and alphabetically
	- Package
	- Imports
	- Constants
	- Variables
	- Types (Alphabetical)
		- New() Functions per Type
		- Bound Functions per Type (Alphabetical)
	- Helper Functions (Alphabetical)
* Helper functions are not exported
* Bound functions may be unexported in the future
* All types and functions documented
	* Full sentences in documentation
* Variables use the shortest possible names while still being descriptive
	* Global variables have longer names
	* Local variables have single-letter names
