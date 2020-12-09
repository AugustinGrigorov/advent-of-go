import org.junit.jupiter.api.Test

class SolutionsTest {
    private val dayOneSolution = day1.Solution()
    private val dayTwoSolution = day2.Solution()
    private val dayThreeSolution = day3.Solution()
    private val dayFourSolution = day4.Solution()

    private val dayOneTestInput = "1721\n" +
            "979\n" +
            "366\n" +
            "299\n" +
            "675\n" +
            "1456"

    @Test
    fun day1_findTwoNumbersCorrectly() {
        assert(dayOneSolution.solve(2, dayOneTestInput) == 514579)
    }

    @Test
    fun day1_findThreeNumbersCorrectly() {
        assert(dayOneSolution.solve(3, dayOneTestInput) == 241861950)
    }

    private val dayTwoTestInput = "1-3 a: abcde\n" +
            "1-3 b: cdefg\n" +
            "2-9 c: ccccccccc"

    @Test
    fun day2_validatesPasswordsWithOldRulesCorrectly() {
        assert(dayTwoSolution.solve(day2.ValidationMethod.OLD, dayTwoTestInput) == 2)
    }

    @Test
    fun day2_validatesPasswordWithNewRulesCorrectly() {
        assert(dayTwoSolution.solve(day2.ValidationMethod.NEW, dayTwoTestInput) == 1)
    }

    private val dayThreeInput = "..##.......\n" +
            "#...#...#..\n" +
            ".#....#..#.\n" +
            "..#.#...#.#\n" +
            ".#...##..#.\n" +
            "..#.##.....\n" +
            ".#.#.#....#\n" +
            ".#........#\n" +
            "#.##...#...\n" +
            "#...##....#\n" +
            ".#..#...#.#"

    @Test
    fun day3_countsTreesCorrectly() {
        assert(dayThreeSolution.solve(dayThreeInput, arrayOf(arrayOf(1, 1))) == 2L)
        assert(dayThreeSolution.solve(dayThreeInput, arrayOf(arrayOf(1, 3))) == 7L)
        assert(dayThreeSolution.solve(dayThreeInput, arrayOf(arrayOf(1, 5))) == 3L)
        assert(dayThreeSolution.solve(dayThreeInput, arrayOf(arrayOf(1, 7))) == 4L)
        assert(dayThreeSolution.solve(dayThreeInput, arrayOf(arrayOf(2, 1))) == 2L)
    }

    @Test
    fun day3_countsTreesOnMultipleTrajectories() {
        assert(
            dayThreeSolution.solve(
                dayThreeInput,
                arrayOf(arrayOf(1, 1), arrayOf(1, 3), arrayOf(1, 5), arrayOf(1, 7), arrayOf(2, 1))
            ) == 336L
        )
    }

    @Test
    fun day4_doesNotCountPassportsWithMissingFields() {
        val passportsWithMissingFields = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\n" +
                "byr:1937 iyr:2017 cid:147 hgt:183cm\n" +
                "\n" +
                "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\n" +
                "hcl:#cfa07d byr:1929\n" +
                "\n" +
                "hcl:#ae17e1 iyr:2013\n" +
                "eyr:2024\n" +
                "ecl:brn pid:760753108 byr:1931\n" +
                "hgt:179cm\n" +
                "\n" +
                "hcl:#cfa07d eyr:2025 pid:166559648\n" +
                "iyr:2011 ecl:brn hgt:59in"
        assert(dayFourSolution.solve(false, passportsWithMissingFields) == 2)
    }

    @Test
    fun day4_doesNotValidateInvalidPassports() {
        val invalidPassportInput = "eyr:1972 cid:100\n" +
                "hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926\n" +
                "\n" +
                "iyr:2019\n" +
                "hcl:#602927 eyr:1967 hgt:170cm\n" +
                "ecl:grn pid:012533040 byr:1946\n" +
                "\n" +
                "hcl:dab227 iyr:2012\n" +
                "ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277\n" +
                "\n" +
                "hgt:59cm ecl:zzz\n" +
                "eyr:2038 hcl:74454a iyr:2023\n" +
                "pid:3556412378 byr:2007"
        assert(dayFourSolution.solve(true, invalidPassportInput) == 0)
    }

    @Test
    fun day4_validatesValidPassports() {
        val validPassports = "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\n" +
                "hcl:#623a2f\n" +
                "\n" +
                "eyr:2029 ecl:blu cid:129 byr:1989\n" +
                "iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm\n" +
                "\n" +
                "hcl:#888785\n" +
                "hgt:164cm byr:2001 iyr:2015 cid:88\n" +
                "pid:545766238 ecl:hzl\n" +
                "eyr:2022\n" +
                "\n" +
                "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"
        assert(dayFourSolution.solve(true, validPassports) == 4)
    }

    @Test
    fun day5_getsCorrectPassportId() {
        val testCases = hashMapOf<String, Int>(
            "BFFFBBFRRR" to 567,
            "FFFBBBFRRR" to 119,
            "BBFFBBFRLL" to 820
        )
        testCases.forEach {
            val dayFiveSolution = day5.Solution(it.key)
            assert(dayFiveSolution.biggest() == it.value)
        }
    }

    val daySixInput = "abc\n" +
            "\n" +
            "a\n" +
            "b\n" +
            "c\n" +
            "\n" +
            "ab\n" +
            "ac\n" +
            "\n" +
            "a\n" +
            "a\n" +
            "a\n" +
            "a\n" +
            "\n" +
            "b"
    private val daySixSolution = day6.Solution(daySixInput)

    @Test
    fun day6_countsAnyAnswersInAGroupCorrectly() {
        assert(daySixSolution.anyone() == 11)
    }

    @Test
    fun day6_countsEveryAnswersInAGroupCorrectly() {
        assert(daySixSolution.everyone() == 6)
    }


    @Test
    fun day7_findsBagParentsCorrectly() {
        val daySevenSolution = day7.Solution(
            "light red bags contain 1 bright white bag, 2 muted yellow bags.\n" +
                    "dark orange bags contain 3 bright white bags, 4 muted yellow bags.\n" +
                    "bright white bags contain 1 shiny gold bag.\n" +
                    "muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.\n" +
                    "shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.\n" +
                    "dark olive bags contain 3 faded blue bags, 4 dotted black bags.\n" +
                    "vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.\n" +
                    "faded blue bags contain no other bags.\n" +
                    "dotted black bags contain no other bags."
        )
        assert(daySevenSolution.parentsOfShinyGoldBag == 4)
    }

    @Test
    fun day7_findsBagChildrenCorrectly() {
        val daySevenSolution = day7.Solution(
            "shiny gold bags contain 2 dark red bags.\n" +
                    "dark red bags contain 2 dark orange bags.\n" +
                    "dark orange bags contain 2 dark yellow bags.\n" +
                    "dark yellow bags contain 2 dark green bags.\n" +
                    "dark green bags contain 2 dark blue bags.\n" +
                    "dark blue bags contain 2 dark violet bags.\n" +
                    "dark violet bags contain no other bags."
        )
        assert(daySevenSolution.childrenOfShinyGoldBag == 126)
    }

    private val dayEightSolution = day8.Solution(
"nop +0\n" +
        "acc +1\n" +
        "jmp +4\n" +
        "acc +3\n" +
        "jmp -3\n" +
        "acc -99\n" +
        "acc +1\n" +
        "jmp -4\n" +
        "acc +6"
    )

    @Test
    fun day8_detectsInfiniteLoopCorrectly() {
        assert(dayEightSolution.run(false) == 5)
    }

    @Test
    fun day8_fixesInfiniteLoopCorrectly() {
        assert(dayEightSolution.run(true) == 8)
    }
}