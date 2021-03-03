package sockets

// Not current as in Reader not go Reader
// func (a *adminSocket) Reader(conn *websocket.Conn) {
//	const op = "AdminSocket.Reader"
//
//	defer conn.Close()
//
//	for {
//		// Read in a message
//		messageType, p, err := conn.ReadMessage()
//		if err != nil {
//			logger.WithError(err)
//			return
//		}
//
//		// Print out that message for clarity
//		logger.Info(string(p))
//
//		err = conn.WriteMessage(messageType, p)
//		if err != nil {
//			logger.Error(err)
//			return
//		}
//	}
//}
