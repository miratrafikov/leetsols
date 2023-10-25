func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    if image[sr][sc] == color {
        return image
    }
    fill(&image, sr, sc, image[sr][sc], color)
    return image
}

func fill(image *[][]int, x, y, startColor, targetColor int) {
    if (*image)[x][y] != startColor {
        return
    }
    (*image)[x][y] = targetColor
    if x-1 >= 0 {
        fill(image, x-1, y, startColor, targetColor)
    }
    if x+1 < len((*image)) {
        fill(image, x+1, y, startColor, targetColor)
    }
    if y-1 >= 0 {
        fill(image, x, y-1, startColor, targetColor)
    }
    if y+1 < len((*image)[0]) {
        fill(image, x, y+1, startColor, targetColor)
    }
}
