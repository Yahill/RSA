# RSA
This is simple RSA encryption.
First you need to generate keys, to do this, write "go run RSA.go -option=keys" .
To crypt message write message into message.txt file than "go run RSA.go -option=crypt -openKey=your open key -signature=your signature" .
You will get crypted message in the crypted_message.txt file.
To decrypt message put crypted information into crypted_message.txt than "go run RSA.go -option=decrypt -closedKey=your closed Key -signature=your signature" .
You will have decrypted message in the decrypted_message.txt.
