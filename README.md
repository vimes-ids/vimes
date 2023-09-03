# vimes

vimes is the best Intrusion Detection System on this side of the Ramptops.
vimes finds unexpected changes in listening ports, files, and kernel modules.

We combine those checks in a single tool to optimize actual security. Any of
those checks by itself is already a good idea, but the combination makes it
really hard for attackers to do something unnoticed.

## Usage

`vimes init` will initialize the database.
`vimes baseline` log the current state into the database.
`vimes check` will check if there are new ports listening, files changes, or kernel modules loaded.

## What's with the name?

vimes is named after His Grace, the Duke of Ankh, Commander Sir Samuel Vimes of
the Ankh-Morpork City Watch, a fictional character in Terry Pratchett's
Discworld series. Vimes is depicted in the novels as somewhere between an
Inspector Morse-type 'old-school' British policeman, and a film noir-esque
grizzled detective.