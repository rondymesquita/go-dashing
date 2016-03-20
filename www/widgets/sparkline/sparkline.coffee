class Dashing.Sparkline extends Dashing.Widget
  displayError:(msg) ->
    $(@node).find(".error").show()
    $(@node).find(".error").html(msg)
  displayMissingDependency:(name,url) ->
    error_html = "<h1>Missing #{name}</h1><p>Download <a href='#{url}'>#{name}</a> and place it in the <span class='highlighted'>assets/javascripts</span> folder"
    @displayError(error_html)

  ready: ->
    @displayMissingDependency("moment.js","http://momentjs.com/downloads/moment.min.js") if (!window.moment)
    @displayMissingDependency("lodash.js","https://raw.githubusercontent.com/lodash/lodash/2.4.1/dist/lodash.min.js") if (!window._)
    @displayMissingDependency("jQuery Sparkline","http://omnipotent.net/jquery.sparkline/#s-about") if (!$.fn.sparkline)

    if @get('debug')
      @debug = (@get('debug'))
    else
      @debug = false

    @isPercent = false
    if @get('is_percent')?
      @isPercent = true

    @prediction = false
    if @get('add_prediction')?
      @prediction = true

    @fixed7Day = 2
    if @get('fixed_7_day')?
      @fixed7Day = @get('fixed_7_day')

    @fixed = 0
    if @get('fixed')?
      @fixed = @get('fixed')
    @diffInverted = @get('diff_inverted')?


    # use the data-unit property in the widget tag to indicate the unit to display (Default:ms)
    if typeof @get('unit') isnt "undefined"
      @unit = (@get('unit'))
    else
      @unit = "ms"

    $n = $(@node)

    self = this

    console.dir self if @debug

  onData: (data) ->
    @renderResults(data)
    @renderSparkline(data)

  renderSparkline: (data) ->
    #dataset = _.compact(roundUpArrayValues(removeTimestampFromTuple(data.series)))
    dataset = removeTimestampFromTuple(data.series)

    if dataset.length>1
      options = {
        type: 'line',
        width:'12em',
        fillColor: false
      }
      if @isPercent
        min = dataset[0] * 100
        max = 100
        hundredDataset = [[0,100], [dataset.length, 100]]
        for _, i in dataset
          dataset[i] *= 100
          min = dataset[i] if dataset[i] < min
          max = dataset[i] if dataset[i] > max
        options['chartRangeMin'] = min
        options['chartRangeMax'] = max
        options['type'] = 'line'
        options['highlightLineColor'] = null
        options['highlightSpotColor'] = null
        options['tooltipFormat'] = ''
        options['lineColor'] = 'gray'
        options['spotColor'] = false
        options['minSpotColor'] = false
        options['maxSpotColor'] = false
        $(@node).find(".sparkline-chart").sparkline(hundredDataset, options)
        delete options['tooltipFormat']
        delete options['highlightLineColor']
        delete options['highlightSpotColor']
        delete options['lineColor']
        delete options['spotColor']
        delete options['minSpotColor']
        delete options['maxSpotColor']
        options['type'] = 'line'

        options['lineWidth'] = 1
        options['composite'] = true
      $(@node).find(".sparkline-chart").sparkline(dataset, options)
    else
      $(@node).find(".sparkline").hide()

  renderResults: (data) ->
    dataAverage = data.series[data.series.length - 1][1]
    change_rate = data.delta
    colors = ["green", "white", "red"]
    if @diffInverted
      colors = ["red", "white", "green"]

    $(@node).find(".change-rate i").removeClass("icon-arrow-up").removeClass("icon-arrow-down")
    if !@isPercent && @prediction && !isNaN(change_rate) && !isNaN(dataAverage) && data.series.length > 7
      dataPrediction = data.series[data.series.length - 8][1] * (1 + change_rate)
      $(@node).find(".prediction").html("(#{dataPrediction.toFixed(@fixed)})")
    if isNaN change_rate
      change_rate = "No data for -7d"
      $(@node).find(".change-rate").css("font-size","1em")
      $(@node).find(".change-rate").css("line-height","40px")

    else
      change_rate = (change_rate * 100).toFixed(@fixed7Day)
      if change_rate>0
        $(@node).find(".change-rate").css("color",colors[0])
        change_rate=change_rate+"%"
        $(@node).find(".change-rate i").addClass("icon-arrow-up")
      else if change_rate==0
        $(@node).find(".change-rate").css("color",colors[1])
        change_rate="no change"
        $(@node).find(".change-rate").css("font-size","1em")
        $(@node).find(".change-rate").css("line-height","40px")
      else
        $(@node).find(".change-rate").css("color",colors[2])
        change_rate=change_rate+"%"
        $(@node).find(".change-rate i").addClass("icon-arrow-down")

    unit = @unit
    if isNaN dataAverage
      $(@node).find(".value").text("N/A")
    else
      if @isPercent
        dataAverage *= 100
      dataAverage = dataAverage.toFixed(@fixed)
      $(@node).find(".value").html("#{dataAverage}<span style='font-size:.3em;'>#{unit}</span>")
    $(@node).find(".change-rate span").text("#{change_rate}")
    $(@node).find(".change-rate span")
    $(@node).find(".updated-at").text(moment().format('MMMM Do YYYY, h:mmA'))

    return

  removeTimestampFromTuple = (arr) ->
    _.map(arr, (num) -> num[1])
  roundUpArrayValues = (arr) ->
    _.map(arr, (num) -> Math.floor(num))

