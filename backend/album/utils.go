package album

// func RemoveElementFromAlbums(id string) bool {
// 	for i, a := range albums {
// 		if a.ID == id {
// 			// Remove element from slice
// 			albums = append(albums[:i], albums[i+1:]...)
// 			return true
// 		}
// 	}

// 	return false
// }

// func AlbumToString(album Album) string {
// 	return fmt.Sprintf("{\"id\":\"%s\",\"title\":\"%s\",\"artist\":\"%s\",\"price\":%.2f}", album.ID, album.Title, album.Artist, album.Price)
// }

// func AlbumSliceToString(objSlice []Album) string {
// 	var sb strings.Builder

// 	sb.WriteString("[")

// 	for i, album := range albums {
// 		sb.WriteString(AlbumToString(album))

// 		if i < len(albums)-1 {
// 			sb.WriteString(",")
// 		}
// 	}

// 	sb.WriteString("]")

// 	return sb.String()
// }
