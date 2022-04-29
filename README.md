"# rate-limiting-with-golang" 

1)
main.go için bir konsol ekranı açılır. 
go run main.go ile sunucu ayağa kaldırılır.

2)
Vegeta için bir konsol ekranı açılır...
vegeta attack -duration=10s -rate=200 -targets=vegeta.conf 
ile "vegeta.conf" dosyası içerisindeki -> " GET http://localhost:8080/ninja "
konumuna saldırı yapılır.

BASİT RATE LİMİTİNG DENEMESİ

This project was created for optimization project. 
KTU COMPUTER ENGINEERING- OPTIMIZATION
