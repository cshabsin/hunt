a = {11, 9, 13, 0, 11, 9, 13, 0, 6, 15, 6}
b = {54, 16, 34, 20, 54}
c = #a
d = #b

e = {}
for f=1,d do
  e[f] = {}
  for g=1,c do
    e[f][g] = 0
  end
end

function h (i)
  if (i == c*d) then
    j = 0
    for k = 1,d do
      l = 0
      for m = 1,c do
        if e[k][m] == 1 then
          l = l + m
        end
      end
      j = j and l == b[k]
    end
    for n = 1,c do
      o = 0
      for p = 1,d do
        if e[p][n] == 1 then
          o = o + p
        end
      end
      j = j and o == a[n]
    end
    if j then
      for q = 1,d do
        r = ""
        for s = 1,c do
          if e[q][s] == 1 then
            r = r .. "X"
          else
            r = r .. "."
          end
        end
        print(r)
      end
    end
  else
    h(i+1)
    t = 1 + i // c
    u = 1 + i % c
    e[t][u] = 1 - e[t][u]
    return h(i + 1)
  end
end

if 0 then
  print("0 is true")
end

print(c*d)

h(0)

for q = 1,d do
  r = ""
  for s = 1,c do
    if e[q][s] == 1 then
      r = r .. "X"
    else
      r = r .. "."
    end
  end
  print(r)
end
