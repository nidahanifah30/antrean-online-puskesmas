document.addEventListener("DOMContentLoaded", function () {
    const searchInput = document.getElementById("searchInput");
    if (!searchInput) return;

    searchInput.addEventListener("keyup", function () {
        const filter = this.value.toLowerCase();
        const rows = document.querySelectorAll("table tbody tr");

        rows.forEach(row => {
        const bulan = row.cells[1]?.textContent.toLowerCase(); // kolom ke-2 = Bulan
        row.style.display = bulan.includes(filter) ? "" : "none";
        });
    });
});
