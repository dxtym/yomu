export type FetchState<T> = 
  | { state: "pending"; } 
  | { state: "resolved"; value: T; } 
  | { state: "rejected"; error: Error }