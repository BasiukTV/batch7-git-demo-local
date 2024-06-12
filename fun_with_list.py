# Some example lists
l1 = [1, 3, 5, 8, 9]
l2 = ["a", "b", "z"]
l3 = list("abrakadabra")

# Printing out the example lists
print(l1, l2, l3)

# Creating lists with list comprehension
l4 = [a * 2 for a in l1]
print(l4)

# list sorting
sl = list(sorted(l3))
print(sl)
