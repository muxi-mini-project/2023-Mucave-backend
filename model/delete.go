package model

func DeleteRelationship(relationship Relationship) error {
	err := DB.Where("followed_id=? AND follower_id= ?", relationship.FollowedId, relationship.FollowerId).Delete(&Relationship{}).Error
	return err
}
