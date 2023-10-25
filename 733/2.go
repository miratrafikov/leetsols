func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    if image[sr][sc] == color {
        return image
    }
    fill(&image, sr, sc, image[sr][sc], color)
    return image
}

func fill(image *[][]int, x, y, startColor, targetColor int) {
    (*image)[x][y] = targetColor
    if x-1 >= 0 && (*image)[x-1][y] == startColor {
        fill(image, x-1, y, startColor, targetColor)
    }
    if x+1 < len((*image)) && (*image)[x+1][y] == startColor {
        fill(image, x+1, y, startColor, targetColor)
    }
    if y-1 >= 0 && (*image)[x][y-1] == startColor{
        fill(image, x, y-1, startColor, targetColor)
    }
    if y+1 < len((*image)[0]) && (*image)[x][y+1] == startColor{
        fill(image, x, y+1, startColor, targetColor)
    }
}
