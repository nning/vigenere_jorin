# Vigenere Jorin

This is a quick POC of an idea of my 6 year old son:

In regular vigenere, you can use Kasiski examination to deduce the key length (and do a frequency analysis).
But what if you change the way, the key position is calculated depending on plain text properties?

We thought of starting on key position 0 on every 2nd or 3rd word change.
This code implements key position reset on 2nd word changes.
(It would also be possible to use rules like "every word change if position is even" or "every word change if previous plain text symbol's index was even").

Of course, this will not change vigenere not quite being fit for after 1863 but we can do it by hand. The autokey attack probably will not work but one drawback is that key length is effectively reduced to about two times the average word length of the language the message is written in. Make sure to at least use padding, key derivation, and message authentication with this... ;)
