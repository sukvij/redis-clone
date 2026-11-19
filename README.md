# redis-clone


l, err := net.Listen("tcp", ":6379")

	conn, err := l.Accept()

	_ = err

	defer conn.Close()

	for {

		buff := make([]byte, 1024)

		conn.Read(buff)

		conn.Write([]byte(string(buff) + " OHK\n"))
	}