/*
https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii/description/

You are given a string word and a non-negative integer k.
Return the total number of substrings of word that contain every vowel ('a', 'e', 'i', 'o', and 'u') at least once and exactly k consonants.

Input: word = "ieaouqqieaouqq", k = 1  Output: 3
Explanation:
The substrings with every vowel and one consonant are:
word[0..5], which is "ieaouq".
word[6..11], which is "qieaou".
word[7..12], which is "ieaouq".
*/
func countOfSubstrings(word string, k int) int64 {
    // Массив для хранения последних позиций каждой гласной
    lastVowelPositions := make([]int, 5)
    lastVowelPositions[0], lastVowelPositions[1], lastVowelPositions[2], lastVowelPositions[3], lastVowelPositions[4] = -1, -1, -1, -1, -1

    consonantCount := 0 // Счетчик согласных
    consonantCountS1, consonantCountS2 := 0, 0 // Счетчики согласных для указателей s1 и s2
    s1, s2 := -1, -1 // Указатели для подсчета подстрок

    var result int64 = 0 // Результат

    for i := 0; i < len(word); i++ {
        // Определяем индекс гласной
        vowelIndex := getVowelIndex(word[i])
        if vowelIndex != -1 {
            lastVowelPositions[vowelIndex] = i // Обновляем последнюю позицию гласной
        } else {
            consonantCount++ // Увеличиваем счетчик согласных
        }

        // Пропускаем, если согласных меньше k
        if consonantCount < k {
            continue
        }

        // Находим минимальную позицию среди всех гласных
        minVowelPosition := min(lastVowelPositions[0], min(lastVowelPositions[1], min(lastVowelPositions[2], min(lastVowelPositions[3], lastVowelPositions[4]))))
        if minVowelPosition == -1 {
            continue // Если хотя бы одна гласная отсутствует, пропускаем
        }

        // Сдвигаем s1, чтобы количество согласных было consonantCount - k
        for s1 < i && consonantCountS1 < consonantCount-k {
            s1++
            if getVowelIndex(word[s1]) == -1 {
                consonantCountS1++
            }
        }

        // Сдвигаем s2, чтобы количество согласных было consonantCount - k + 1
        for s2 < i && consonantCountS2 < consonantCount-k+1 {
            s2++
            if getVowelIndex(word[s2]) == -1 {
                consonantCountS2++
            }
        }

        // Вычисляем количество подходящих подстрок
        validSubstrings := max(min(minVowelPosition, s2) - s1, 0)
        result += int64(validSubstrings)
    }

    return result
}

// Функция для определения индекса гласной
func getVowelIndex(c byte) int {
    switch c {
    case 'a':
        return 0
    case 'e':
        return 1
    case 'i':
        return 2
    case 'o':
        return 3
    case 'u':
        return 4
    default:
        return -1
    }
}

// Вспомогательные функции для min и max
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
