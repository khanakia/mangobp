package api_fn

// func IfUserSa(ctx context.Context) bool {
// 	currentUser, err := middleware.GetUserFromGqlContext(ctx)

// 	if err != nil || (currentUser.RoleID != auth_type.RoleSaID) {
// 		return false
// 	}

// 	return true
// }

// func IfUserSaWithError(ctx context.Context) error {
// 	if !IfUserSa(ctx) {
// 		return errors.New("access denied")
// 	}

// 	return nil
// }

// func IfUserSaOrMember(ctx context.Context) bool {
// 	currentUser, err := middleware.GetUserFromGqlContext(ctx)

// 	// goutil.PrintToJSON(currentUser)
// 	// only SA and Member are allowed in future we can have more user roles
// 	if err != nil || (currentUser.RoleID != auth_type.RoleSaID && currentUser.RoleID != auth_type.RoleMemberID) {
// 		return false
// 	}

// 	return true
// }

// func IfUserSaOrMemberWithError(ctx context.Context) error {
// 	if !IfUserSaOrMember(ctx) {
// 		return errors.New("access denied")
// 	}

// 	return nil
// }

// // if user is not SA then user can only see his own order
// func IfUserSaOrRecordOwner(recordUserID string, ctx context.Context) bool {
// 	currentUser, _ := middleware.GetUserFromGqlContext(ctx)

// 	if currentUser.RoleID != auth_type.RoleSaID && recordUserID != currentUser.ID {
// 		return false
// 	}

// 	return true
// }

// func IfUserSaOrRecordOwnerWithError(recordUserID string, ctx context.Context) error {
// 	if !IfUserSaOrRecordOwner(recordUserID, ctx) {
// 		return errors.New("access denied")
// 	}
// 	return nil
// }
