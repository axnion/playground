"use strict"

const router    = require("express").Router()
const mongoose  = require("mongoose")

let uri = "mongodb://localhost:27017/entries"
let connection = mongoose.createConnection(uri)

let EntrySchema = new mongoose.Schema({
    text: {type: String, required: true}
})
let Entry = connection.model("Entry", EntrySchema)

router.route("/").get(function(req, res, next) {
    get().then(function(entries) {
        console.log("Fetched entries")
        res.render("home", {entries: entries})
    }).catch(function(err) {
        console.log(err)
    })
})

router.route("/add").post(function(req, res, next) {
    let text = req.body.text

    add(text).then(function() {
        console.log("Added entry: " + text)
        res.redirect("/")
    }).catch(function(err) {
        console.log(err)
    })
})

router.route("/display").get(function(req, res) {
    let text = req.query.text

    res.render("display", {text: text})
})

function add(text) {
    return new Promise(function(resolve, reject) {
        let entry = new Entry({
            text: text
        })

        entry.save(function(err) {
            if(err) {
                reject(err)
                return
            } else {
                resolve()
            }
        })
    })
}

function get() {
    return new Promise(function(resolve, reject) {
        Entry.find({}, function(err, entries) {
            if (err) {
                reject(err)
            }

            resolve(entries)
        })
    })
}

module.exports = router
