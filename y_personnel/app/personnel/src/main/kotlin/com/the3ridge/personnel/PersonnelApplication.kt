package com.the3ridge.personnel

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class PersonnelApplication

fun main(args: Array<String>) {
    runApplication<PersonnelApplication>(*args)
}
