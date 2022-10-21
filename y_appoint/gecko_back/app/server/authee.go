package server

//
//func (a *AuthServer) Login(ctx context.Context, req *authentication.LoginRequest) (*authentication.LoginResponse, error) {
//	crossLogging.Logger.Info("login request", zap.String("username", req.Username))
//
//	var conn dbConn.Conn
//	db, err := conn.PassConn()
//	if err != nil {
//		crossLogging.Logger.Error("error while creating a new auth repo", zap.Error(err))
//		return nil, err
//	}
//
//	user, err := dbConn.FetchUser(ctx, db, req.Username)
//	if err != nil {
//		crossLogging.Logger.Error("error while logging in", zap.Error(err))
//		return nil, err
//	}
//
//	return &authentication.LoginResponse{
//		Token: user.Token,
//	}, nil
//
//}
