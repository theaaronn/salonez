document.addEventListener('htmx:configRequest', function(event) {
    const userID = localStorage.getItem('userID');
    const userType = localStorage.getItem('userType');
    if (userID && userType) {
        event.detail.headers['X-User-ID'] = userID;
        event.detail.headers['X-User-Type'] = userType;
    }
});
