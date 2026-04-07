async function loadRooms(){

    const res = await fetch("/rooms")

    const rooms = await res.json()

    const container = document.getElementById("rooms")

    container.innerHTML=""

    rooms.forEach(room=>{

        container.innerHTML += `

<div class="room">

<h3>${room.name}</h3>

<p>📍 ${room.location}</p>

<p>👥 Вместимость: ${room.capacity}</p>

<button onclick="bookRoom(${room.id})">Забронировать</button>

</div>

`

    })

}

function bookRoom(id){

    alert("Форма бронирования будет здесь. Комната: "+id)

}
async function loadRooms(){

    const res = await fetch("/rooms")
    const data = await res.json()

    const div = document.getElementById("rooms")

    div.innerHTML=""

    data.forEach(r=>{
        div.innerHTML += `
<div class="card">
<h3>${r.name}</h3>
<p>${r.location}</p>
<p>${r.capacity}</p>
</div>
`
    })

}

async function createBooking(e){

    e.preventDefault()

    const booking = {
        room_id: parseInt(document.getElementById("room_id").value),
        user_id: 1,
        start_time: document.getElementById("start").value,
        end_time: document.getElementById("end").value,
        purpose: document.getElementById("purpose").value
    }

    await fetch("/bookings",{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify(booking)
    })

    alert("Бронь создана")
}