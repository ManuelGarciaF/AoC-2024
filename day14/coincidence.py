#!/usr/bin/env python3

# Numbers taken from looking at the weird strips in the images
firstH = 75
offsetH = 103
firstV = 3
offsetV = 101

hs = [firstH + h for h in range(0, 15000, offsetH)]
vs = [firstV + v for v in range(0, 15000, offsetV)]

intersection = set(hs) & set(vs)
print(min(intersection))
