class Solution {
public:
    int numBusesToDestination(vector<vector<int>>& routes, int source, int target) {
        if (source == target) {
            return 0;
        }

        // Создание карты, где ключ - номер остановки, значение - маршруты, проходящие через эту остановку
        map<int, vector<int>> adjacencyList;
        for (int route = 0; route < routes.size(); route++) {
            for (auto stop : routes[route]) {
                adjacencyList[stop].push_back(route);
            }
        }

        queue<int> queue;
        set<int> visitedRoutes;
        // Добавление в очередь все маршруты, проходящие через исходную остановку
        for (auto route : adjacencyList[source]) {
            queue.push(route);
            visitedRoutes.insert(route);
        }

        int busCount = 1;
        while (!queue.empty()) {
            int size = queue.size();
            for (int i = 0; i < size; i++) {
                int currentRoute = queue.front();
                queue.pop();

                // Перебор остановок в текущем маршруте
                for (auto stop : routes[currentRoute]) {
                    // Возвращение текущего счетчика, если достигнута целевая остановка
                    if (stop == target) {
                        return busCount;
                    }

                    // Перебор следующих возможных маршрутов из текущей остановки
                    for (auto nextRoute : adjacencyList[stop]) {
                        if (visitedRoutes.count(nextRoute) == 0) {
                            visitedRoutes.insert(nextRoute);
                            queue.push(nextRoute);
                        }
                    }
                }
            }
            busCount++;
        }
        return -1;
    }
};
