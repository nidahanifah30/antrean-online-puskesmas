document.addEventListener('DOMContentLoaded', function () {
    const modal = document.getElementById('modalHapus');
    modal.addEventListener('show.bs.modal', function (event) {
        const button = event.relatedTarget;
        const tahun = button.getAttribute('data-tahun');
        const bulan = button.getAttribute('data-bulan');

        document.getElementById('hapus-tahun').innerText = tahun;
        document.getElementById('hapus-bulan').innerText = bulan;

        const form = document.getElementById('formHapus');
        form.action = `/petugas/riwayat/${tahun}/${bulan}/hapus`;
    });
});
