# Array
  Fixed size
  Copy on assignment
  Function gets a full copy
  a := [3]int{1,2,3}
  b := a
  b[0] = 100
  // a unchanged

# Slice
  Has pointer + len + cap
  Assignment shares same array
  a := []int{1,2,3}
  b := a
  
  b[0] = 100
  // a also changes

# append rule (MOST IMPORTANT)
  b := append(a, x)
  Two cases:
  ✅ If capacity is enough → same array (shared)
  ❌ If capacity is full → new array (not shared)
