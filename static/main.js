angular.module('goshort', [])
    .controller('GoShortController', function($http) {
        var goshort = this;

        goshort.shortenUrl = function() {
            $http.post('/create', { url: goshort.url})
                .then(function(data) {
                    console.log('success')
                    console.log(data.data.ShortCode)
                    goshort.shortUrl = data.data.ShortCode
                }, function() {
                    console.log('error')
                })
        }
    })

