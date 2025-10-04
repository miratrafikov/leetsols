func spellchecker(wordlist []string, queries []string) []string {
    n := len(wordlist)
    m := len(queries)

    // asIs is just wordlist turned into a map for quick retrieval for precedence case #1.
    asIs := make(map[string]struct{}, n)
    for _, word := range wordlist {
        asIs[word] = struct{}{}
    }

    // cap is used to retrieve first match for precedence case #2.
    cap := make(map[string]string, n)
    for _, word := range wordlist {
        lowercasedWord := strings.ToLower(word)
        if _, found := cap[lowercasedWord]; !found {
            cap[lowercasedWord] = word
        }
    }

    // vow is used to retrieve first match for precedence case #3.
    vow := make(map[string]string, n)
    for _, word := range wordlist {
        devowelizedWord := devowelize(word)
        if _, found := vow[devowelizedWord]; !found {
            vow[devowelizedWord] = word
        }
    }

    result := make([]string, 0, m)
    for _, qw := range queries {
        if _, found := asIs[qw]; found {
            result = append(result, qw)
            continue
        }

        lowercasedQw := strings.ToLower(qw)
        if match, found := cap[lowercasedQw]; found {
            result = append(result, match)
            continue
        }

        devowelizedQw := devowelize(qw)
        if match, found := vow[devowelizedQw]; found {
            result = append(result, match)
            continue
        }

        result = append(result, "")
    }
    return result
}

func devowelize(s string) string {
    lowercased := strings.ToLower(s)
    var sb strings.Builder
    for i := 0; i < len(lowercased); i++ {
        if lowercased[i] == 'a' || lowercased[i] == 'e' || 
            lowercased[i] == 'i' || lowercased[i] == 'o' || lowercased[i] == 'u' {
            sb.WriteByte('_')
        } else {
            sb.WriteString(lowercased[i:i+1])
        }
    }
    return sb.String()
}
