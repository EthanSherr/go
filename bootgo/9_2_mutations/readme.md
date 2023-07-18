# MUTATIONS
## INSERT AN ELEMENT
```go
m[key] = elem
```

## GET AN ELEMENT
```go
elem = m[key]
```


## DELETE AN ELEMENT
```go
delete(m, key)
```

## CHECK IF A KEY EXISTS
```go
elem, ok := m[key]
```
If `key` is in `m`, then `ok` is `true`. If not, `ok` is `false`.

If `key` is not in the map, then `elem` is the zero value for the map's element type.

## ASSIGNMENT
It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

Complete the `deleteIfNecessary` function.

- If the user doesn't exist in the map, return the error `not found`.
- If they exist but aren't scheduled for deletion, return `deleted` as `false` with no errors.
- If they exist and are scheduled for deletion, return `deleted` as `true` with no errors and delete their record from the map.

## NOTE ON PASSING MAPS
Like slices, maps are also passed by reference into functions. This means that when a map is passed into a function we write, we can make changes to the original, we don't have a copy.