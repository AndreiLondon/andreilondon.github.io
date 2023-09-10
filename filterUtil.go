package main

func filterByCategories(posts []Post, filterCategories []string) []Post {
	if len(filterCategories) == 0 {
		return posts
	}
	result := []Post{}
	for _, post := range posts {
		categoties := post.Categories
		if containsArr(filterCategories, categoties) {
			result = append(result, post)
		}
	}
	return result
}
func filterByMode(posts []Post, mode string, user *User) []Post {
	if user == nil {
		return posts
	}
	if mode == SHAW_ALL {
		return posts
	}
	switch mode {
	case MY_POSTS:
		filteresPosts := []Post{}
		for _, post := range posts {
			if post.Username == user.Username {
				filteresPosts = append(filteresPosts, post)
			}
		}
		return filteresPosts
	case MY_COMMENTS:
		filteresPosts := []Post{}
		for _, post := range posts {
			comments := post.Comments
			for _, comment := range comments {
				if comment.Username == user.Username {
					filteresPosts = append(filteresPosts, post)
				}
			}
		}
		return filteresPosts
	case MY_LIKES:
		filteresPosts := []Post{}
		for _, post := range posts {
			if post.Status != 0 {
				filteresPosts = append(filteresPosts, post)
				continue
			}
			comments := post.Comments
			for _, comment := range comments {
				if comment.Status != 0 {
					filteresPosts = append(filteresPosts, post)
					break
				}
			}
		}
		return filteresPosts
	default:
		return posts
	}
}
