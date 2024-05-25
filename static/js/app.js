document.addEventListener('DOMContentLoaded', function () {
    //menu dropdown
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

    //tabs active and not active
    const tabs = document.querySelectorAll('.tabs li a');
    // Add click event listener to each 'a' element
    // Function to remove the 'active' class from all elements
    function removeActiveClasses() {
        tabs.forEach(tab => {
            tab.classList.remove('active');
        });
    }
    tabs.forEach(tab => {
        tab.addEventListener('click', function () {
            // Remove 'active' class from all 'a' elements
            removeActiveClasses();
            // Add 'active' class to the clicked 'a' element
            this.classList.add('active');
        });
    });
});


document.addEventListener("DOMContentLoaded", function () {
    const tabs = document.querySelectorAll(".normaltabs a");
    const contents = document.querySelectorAll(".tab-content");

    tabs.forEach(tab => {
        console.log(tab);
        tab.addEventListener("click", function (event) {
            event.preventDefault();

            // Remove active class from all tabs and hide all content
            tabs.forEach(t => t.classList.remove("active"));
            contents.forEach(content => content.classList.add("hidden"));

            // Add active class to the clicked tab and show the corresponding content
            this.classList.add("active");
            const activeContent = document.getElementById(this.dataset.tab);
            activeContent.classList.remove("hidden");
        });
    });
});
