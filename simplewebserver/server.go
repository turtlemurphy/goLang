package main

/* func main() {
	//http.ListenAndServe("", nil)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}

type Server struct {
	Addr           string
	Handler        Handler
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
	TLSConfig      *tls.Config
	TLSNextProto   map[string]func(*Server, *tls.Conn, Handler)
	ConnState      func(net.Conn, ConnState)
	ErrorLog       *log.Logger
}
*/

func main() {
	// //	max := new(big.Int).Lsh(big.NewInt(1), 128)
	// 	//serial Number, _ := rand.Int(rand.Reader, max)
	// 	subject := pkix.Name{
	// 		Organization:		[]string{"Manning Publications Co."},
	// 		OrganizationalUnit:	[]string{"Books"},
	// 		CommonName:			"Go Web Programming",
	// 	}
	// 	template := x509.Certificate{
	// 		SerialNumber:	serialNumber,
	// 		Subject:		subject,
	// 		NotBefore:		time.Now(),
	// 		NotAfter:		time.Now().Add(365 * 24 * time.Hour),
	// 		KeyUsage:		x509.KeyUsageKeyEncipherment	|	x509.KeyUsageDigitalSignature,
	// 		ExtKeyUsage:	[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	// 		IPAddresses:	[]net.IP{net.ParseIP("127.0.0.1")},
	// 	}
	// 	pk, _ := rsa.GenerateKey(rand.Reader, 2048)
	// 	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	// 	cartOut, _ := os.Create("cert.pem")
	// 	pem.Encode(certOut, &pem.Block{Type:	"CERTIFICATE", Bytes:derBytes})
	// 	certOut.Close()
	// 	keyOut, _ := os.Create("key.pem")
	// 	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	// 	keyOut.Close()
}
