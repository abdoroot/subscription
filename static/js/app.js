document.addEventListener('DOMContentLoaded', function () {
    //menu dropdown
    document.body.addEventListener('click', function (event) {
        const button = event.target.closest('[data-collapse-toggle]');
        if (button) {
            const menuId = button.getAttribute('data-collapse-toggle');
            const menu = document.getElementById(menuId);
            if (menu) {
                if (menu.classList.contains('slide-down')) {
                    menu.classList.remove('slide-down');
                    menu.classList.add('slide-up');
                } else {
                    menu.classList.remove('slide-up');
                    menu.classList.add('slide-down');
                }
            }
        }
    });

    //tabs active and not active
    function removeActiveClasses(tabs) {
        tabs.forEach(tab => {
            tab.classList.remove('active');
        });
    }

    document.body.addEventListener('click', function (event) {
        const tab = event.target.closest('.tabs li a');
        if (tab) {
            const tabs = document.querySelectorAll('.tabs li a');
            removeActiveClasses(tabs);
            tab.classList.add('active');
        }
    });

    document.body.addEventListener('click', function (event) {
        const tab = event.target.closest('.normaltabs a');
        if (tab) {
            event.preventDefault();
            const tabs = document.querySelectorAll('.normaltabs a');
            const contents = document.querySelectorAll('.tab-content');

            tabs.forEach(t => t.classList.remove('active'));
            contents.forEach(content => content.classList.add('hidden'));

            tab.classList.add('active');
            const activeContent = document.getElementById(tab.dataset.tab);
            if (activeContent) {
                activeContent.classList.remove('hidden');
            }
        }
    });

    //Modals
    const overlayDiv = document.getElementById("overlay");

    document.body.addEventListener('click', function (event) {
        const btn = event.target.closest('[data-target-modal]');
        if (btn) {
            const modalId = btn.getAttribute('data-target-modal');
            const modalElement = document.getElementById(modalId);
            if (modalElement) {
                modalElement.classList.toggle('hidden');
                overlayDiv.classList.toggle('hidden');
            } else {
                console.error(`No element found with ID ${modalId}`);
            }
        }

        const closeModalBtn = event.target.closest('.close-modal');
        if (closeModalBtn) {
            const modals = document.querySelectorAll('.modal');
            modals.forEach(modal => {
                modal.classList.add('hidden');
            });
            overlayDiv.classList.add('hidden');
        }
    });
});
