package database

type Video struct {
	VideoId    int    `json:"video_id"`
	AuthorId   int    `json:"author_id"`
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
	CreateTime string `json:"create_time"`
}

func (v *Video) Add() error {
	input, err := db.Prepare("INSERT INTO `videos` (author_id,name,category_id,create_time) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = input.Exec(v.AuthorId, v.Name, v.CategoryId, v.CreateTime)
	defer input.Close()
	return err
}

func (v *Video) GetByVideoId() (*Video, error) {
	output, err := db.Prepare("SELECT * FROM `videos` WHERE video_id=?")
	if err != nil {
		return nil, err
	}
	var video Video
	err = output.QueryRow(v.VideoId).Scan(&video.VideoId, &video.AuthorId, &video.Name, &video.CategoryId, &video.CreateTime)
	if err != nil {
		return nil, err
	}
	defer output.Close()
	return &video, nil
}

func (v *Video) GetByAuthorId() ([]Video, error) {
	output, err := db.Prepare("SELECT * FROM `videos` WHERE author_id=?")
	if err != nil {
		return nil, err
	}
	rows, err := output.Query(v.AuthorId)
	if err != nil {
		return nil, err
	}
	var videos []Video
	for rows.Next() {
		var video Video
		if err := rows.Scan(&video.VideoId, &video.AuthorId, &video.Name, &video.CategoryId, &video.CreateTime); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	defer output.Close()
	return videos, nil
}

func (v *Video) Delete() error {
	input, err := db.Prepare("DELETE FROM `videos` WHERE video_id=? AND author_id=?")
	if err != nil {
		return err
	}
	_, err = input.Exec(v.VideoId, v.AuthorId)
	return err
}
