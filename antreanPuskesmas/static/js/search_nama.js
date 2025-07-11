document.addEventListener("DOMContentLoaded", function () {
    const searchInput = document.getElementById("searchInput");

    if (!searchInput) return;

    searchInput.addEventListener("keyup", function () {
        const filter = this.value.toLowerCase();
        const rows = document.querySelectorAll("table tbody tr");

        rows.forEach(row => {
            const namaCell = row.querySelector(".nama-pasien");
            if (namaCell) {
                const namaText = namaCell.textContent.toLowerCase();
                row.style.display = namaText.includes(filter) ? "" : "none";
            }
        });
    });
});
