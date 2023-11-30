func compress(chars []byte) int {
    encodingEnd := 0
    sequenceStart := 0
    sequenceEnd := 0
    for sequenceEnd < len(chars) {
        if chars[sequenceEnd] == chars[sequenceStart] {
            sequenceEnd++
            if sequenceEnd != len(chars) {
                continue
            }
        }
        chars[encodingEnd] = chars[sequenceStart]
        encodingEnd++
        sequenceLength := sequenceEnd - sequenceStart
        if sequenceLength > 1 {
            lengthToChars := []byte(strconv.Itoa(sequenceLength))
            for i := 0; i < len(lengthToChars); i, encodingEnd = i + 1, encodingEnd + 1 {
                chars[encodingEnd] = lengthToChars[i]
            }
        }
        sequenceStart = sequenceEnd
    }
    return encodingEnd
}
