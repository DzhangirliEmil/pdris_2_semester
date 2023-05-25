# pdris_2_semester
В кластере используется два сервиса - nginx и jokeservice.

nginx перенаправляет внутренний трафик.

jokeservice ходит на внешнюю ручку https://geek-jokes.sameerkumar.website/api?format=json, которая возвращает случайную шутку :)

Реализацию сервиса jokeservice можно посмотреть в директории jokeservice 

Трафик в кластер приходит через ingress-gateway, выходит в внешнюю сеть через egress-gateway.

Для запуска необходимо:
1. Запустить скрипт run.sh, не закрывать заблокированный терминал.
2. В другом терминале запустить скрипт kubectl get svc -n istio-system istio-ingressgateway, посмотреть External Ip сервиса
3. Открыть в браузере вкладку http://$x/joke, где x - External Ip сервиса
