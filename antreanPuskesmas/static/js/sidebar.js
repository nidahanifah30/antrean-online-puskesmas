document.addEventListener("DOMContentLoaded", function () {
    const toggleBtn = document.getElementById("toggleSidebar");
    const sidebar = document.getElementById("sidebar");
    const mainContent = document.getElementById("mainContent");

    toggleBtn?.addEventListener("click", function () {
        sidebar.classList.toggle("hidden");
        mainContent.classList.toggle("full");
    });
});
