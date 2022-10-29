package com.the3ridge.personnel.routes.home

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/home")
class HomeController {

    @GetMapping("{id}")
    fun getPersonnelById(@PathVariable id: Long): Long {
        return id
    }

}