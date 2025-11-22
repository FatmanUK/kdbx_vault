# kdbx_vault

A networked KDBX interface on the command line. A cross between Hashicorp Vault's API and KeepAss2/KeepAssXC.

A wrapper around this library:
https://github.com/tobischo/gokeepasslib

Note: by its nature, this wants a password in the clear over HTTP. It's designed to be used on the command line for localhost only, however it is technically possible to put it online. If you're putting it online, for the love of cheese put it behind a reverse TLS proxy on a trusted server. The password will still be in plain text between the proxy and kdbx_vault, so you must be sure the proxy doesn't log headers. Oh, let's do a base64 encode, anyway.

Questions I asked myself in lieu of FAQ (it's useful for me to remember this stuff):

Q: Will I be adding creation of KDBX databases?

A: Not unless I get very bored one weekend. No, you can create them using one of the standard tools KeepAss2 or KeepAssXC.

Q: Will I be adding deletion of KDBX databases?

A: No, operating systems already have commands to delete files.
