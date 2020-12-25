### Coding

- Fyne APIs for dropdown works like a charm. I got totally shocked by the data formats in Golang (called layout).
- Pointers got me into some trouble as some initial state appear to leave them as nil.
- Resorting to global state. Not sure this is appropriate, but it makes the logic simple to implement, it probably isn't ideal for unit testing though as global state is likely problematic.
- Returning errors, go won't let a function only return one type if it is declared to return more, and won't allow returning nil for the case of Time type, and it won't allow inline operation on func returned typed it seems.
- The _ declaration is interesting to say the least. Useful as it makes the invoked function signature very explicit.
  

### Issues

- No red colored field, fyne simply does not support it, and has no intention to allow such control over styling.

### Timing

- About 3h. 
