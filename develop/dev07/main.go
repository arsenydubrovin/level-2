package main

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
- cd <args> - смена директории
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента
- ps - выводит общую информацию по запущенным процессам в формате
*/

func main() {
	runShell()
}