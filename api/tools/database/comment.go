package database

type Comment struct {
	CommentId  int    `json:"comment_id"`
	UserId     int    `json:"user_id"`
	Username   string `json:"username"`
	VideoId    int    `json:"video_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

func (c *Comment) Add() error {
	input, err := db.Prepare("INSERT INTO `comments` (user_id,video_id,content,create_time) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = input.Exec(c.UserId, c.VideoId, c.Content, c.CreateTime)
	if err != nil {
		return err
	}
	return nil
}

func GetComments(videoId int) ([]Comment, error) {
	output, err := db.Prepare(`
		SELECT u.username,c.content,c.create_time FROM comments AS c
		LEFT JOIN users AS u on u.user_id=c.user_id
		WHERE video_id=? ORDER BY comment_id DESC`)
	if err != nil {
		return nil, err
	}
	rows, err := output.Query(videoId)
	if err != nil {
		return nil, err
	}
	var comments []Comment
	for rows.Next() {
		var c Comment
		if err := rows.Scan(&c.Username, &c.Content, &c.CreateTime); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	if err := rows.Err(); err != nil || len(comments) == 0 {
		return nil, err
	}
	return comments, nil
}
