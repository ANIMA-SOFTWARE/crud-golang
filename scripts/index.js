/// <reference lib="dom" />
/// <reference lib="dom.iterable" />

//test 5

function toggleTheme() {
    document.documentElement.className = (document.documentElement.className == 'light') ? 'dark' : 'light';
}

let selectedTab
function selectTab(tabButton) {
    selectedTab?.removeAttribute('id')

    tabButton.setAttribute('id','selectedTab')

    selectedTab = tabButton
}

function fillTemplate() {

}
