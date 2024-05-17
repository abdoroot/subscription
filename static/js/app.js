//menu dropdown
document.addEventListener('DOMContentLoaded', function () {
    const buttons = document.querySelectorAll('[data-collapse-toggle]');
    buttons.forEach(function (button) {
        const menuId = button.getAttribute('data-collapse-toggle');
        const menu = document.getElementById(menuId);

        button.addEventListener('click', function () {
            if (menu.classList.contains('slide-down')) {
                menu.classList.remove('slide-down');
                menu.classList.add('slide-up');
            } else {
                menu.classList.remove('slide-up');
                menu.classList.add('slide-down');
            }
        });
    });
});

