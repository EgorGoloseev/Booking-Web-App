async function loadMyBookings() {
    const res = await fetch("/bookings/my?from=2025-01-01&to=2030-12-31", {
        headers: { "X-User-Id": 1 }
    });
    const bookings = await res.json();
    const tbody = document.getElementById("myBookings");
    tbody.innerHTML = "";

    bookings.forEach(b => {
        tbody.innerHTML += `
      <tr>
        <td>${b.room_id}</td>
        <td>${new Date(b.start_time).toLocaleString()}</td>
        <td>${new Date(b.end_time).toLocaleString()}</td>
        <td>${b.purpose}</td>
        <td><button onclick="cancelBooking(${b.id})" class="btn">Отменить</button></td>
      </tr>
    `;
    });
}

async function cancelBooking(id) {
    const res = await fetch(`/bookings/${id}`, {
        method: "DELETE",
        headers: { "X-User-Id": 1 }
    });
    if (res.status === 200) loadMyBookings();
    else alert("Ошибка при удалении брони");
}

loadMyBookings();