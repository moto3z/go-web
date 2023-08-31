#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <algorithm>

namespace name
{
    
} // namespace name


std::map<std::string, int> getPointMapper(const std::string &inputString);
void getSubStringRecursive(int beginIndex, const std::string &inputString, std::set<std::string> &subStrings);

int main() {
    int n;
    std::cin >> n;
    std::string inputString;
    std::cin >> inputString;

    std::map<std::string, int> pointMapper = getPointMapper(inputString);
    int maxi = -1;

    for (int first = 1; first < inputString.length(); ++first) {
        for (int second = first + 1; second < inputString.length(); ++second) {
            std::string section1 = inputString.substr(0, first);
            std::string section2 = inputString.substr(first, second - first);
            std::string section3 = inputString.substr(second, inputString.length() - second);

            maxi = std::max(
                pointMapper[section1] + pointMapper[section2] + pointMapper[section3],
                maxi
            );
        }
    }
    std::cout << maxi + 3 << std::endl;

    return 0;
}

std::map<std::string, int> getPointMapper(const std::string &inputString) {
    std::set<std::string> subStrings;
    getSubStringRecursive(0, inputString, subStrings);

    std::vector<std::string> stringList(subStrings.begin(), subStrings.end());
    std::sort(stringList.begin(), stringList.end());

    std::map<std::string, int> result;
    for (size_t i = 0; i < stringList.size(); ++i) {
        result[stringList[i]] = i;
    }
    return result;
}

void getSubStringRecursive(int beginIndex, const std::string &inputString, std::set<std::string> &subStrings) {
    if (beginIndex >= inputString.length()) {
        return;
    }
    for (int step = 1; step < inputString.length() - beginIndex; ++step) {
        if (beginIndex + step > inputString.length()) {
            return;
        }
        subStrings.insert(inputString.substr(beginIndex, step));
        getSubStringRecursive(beginIndex + step, inputString, subStrings);
    }
}
