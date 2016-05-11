package sqrape

var (
	// twitter, y u no clean up ur html
	tweetRawHTML = `<li class="js-stream-item stream-item stream-item expanding-stream-item" data-item-id="641830475854536704" id="stream-item-tweet-641830475854536704" data-item-type="tweet">
      <div class="tweet original-tweet js-original-tweet js-stream-tweet js-actionable-tweet js-profile-popup-actionable with-non-tweet-action-follow-button" data-tweet-id="641830475854536704" data-disclosure-type="" data-item-id="641830475854536704" data-permalink-path="/nuts4ag/status/641830475854536704" data-screen-name="nuts4ag" data-name="Sara Savary" data-user-id="29542534" data-is-reply-to="true" data-has-parent-tweet="true" data-expanded-footer="<div class=&quot;js-tweet-details-fixer tweet-details-fixer&quot;>
            <div class=&quot;entities-media-container js-media-container&quot; style=&quot;min-height:0px&quot;>
            </div>
      <div class=&quot;js-machine-translated-tweet-container&quot;></div>
        <div class=&quot;js-tweet-stats-container tweet-stats-container &quot;>
        </div>
      <div class=&quot;client-and-actions&quot;>
      <span class=&quot;metadata&quot;>
        <span>5:28 AM - 10 Sep 2015</span>
           &amp;middot; <a class=&quot;permalink-link js-permalink js-nav&quot; href=&quot;/nuts4ag/status/641830475854536704&quot;  tabindex=&quot;-1&quot;>Details</a>
      </span>
    </div>
    </div>
    " data-mentions="HomeDepot" data-you-follow="false" data-follows-you="false" data-you-block="false">
        <div class="context">
        </div>
        <div class="content">
          <div class="stream-item-header">
              <a class="account-group js-account-group js-action-profile js-user-profile-link js-nav" href="/nuts4ag" data-user-id="29542534">
        <img class="avatar js-action-profile-avatar" src="https://pbs.twimg.com/profile_images/630167157687451648/_9Qh1lqB_bigger.jpg" alt="">
        <strong class="fullname js-action-profile-name show-popup-with-id" data-aria-label-part="">Sara Savary</strong>
        <span>‏</span><span class="username js-action-profile-name" data-aria-label-part=""><s>@</s><b>nuts4ag</b></span>
      </a>
            <small class="time">
      <a href="/nuts4ag/status/641830475854536704" class="tweet-timestamp js-permalink js-nav js-tooltip" title="5:28 AM - 10 Sep 2015"><span class="_timestamp js-short-timestamp js-relative-timestamp" data-time="1441859284" data-time-ms="1441859284000" data-long-form="true" aria-hidden="true">16h</span><span class="u-hiddenVisually" data-aria-label-part="last">16 hours ago</span></a>
    </small>
          </div>
            <p class="TweetTextSize  js-tweet-text tweet-text" data-aria-label-part="0" lang="en"><a href="/HomeDepot" class="twitter-atreply pretty-link js-nav" dir="ltr"><s>@</s><b>HomeDepot</b></a> best talk of the UCDavis <a href="/hashtag/bee?src=hash" data-query-source="hashtag_click" class="twitter-hashtag pretty-link js-nav" dir="ltr"><s>#</s><b>bee</b></a> and <a href="/hashtag/Neonic?src=hash" data-query-source="hashtag_click" class="twitter-hashtag pretty-link js-nav" dir="ltr"><s>#</s><b><strong>Neonic</strong></b></a> conference was Ray Jarvis. Doing good work with your sustainability program</p>
      <div class="expanded-content js-tweet-details-dropdown">
      </div>
          <div class="stream-item-footer">
    <a class="details with-icn js-details " href="/nuts4ag/status/641830475854536704">
            <span class="Icon Icon--conversation"></span>
      <b>
          <span class="expand-stream-item js-view-details">View conversation</span>
          <span class="collapse-stream-item js-hide-details">Hide conversation</span>
      </b>
    </a>
          <span class="ProfileTweet-action--reply u-hiddenVisually"></span>
          <span class="ProfileTweet-action--retweet u-hiddenVisually">
            <span class="ProfileTweet-actionCount" aria-hidden="true" data-tweet-stat-count="0">
              <span class="ProfileTweet-actionCountForAria">0 retweets</span>
            </span>
          </span>
          <span class="ProfileTweet-action--favorite u-hiddenVisually">
            <span class="ProfileTweet-actionCount" aria-hidden="true" data-tweet-stat-count="0">
              <span class="ProfileTweet-actionCountForAria">0 favorites</span>
            </span>
          </span>
        <div role="group" aria-label="Tweet actions" class="ProfileTweet-actionList u-cf js-actions ">
          <div class="ProfileTweet-action ProfileTweet-action--reply">
            <button class="ProfileTweet-actionButton u-textUserColorHover js-actionButton js-actionReply" data-modal="ProfileTweet-reply" type="button">
              <div class="IconContainer js-tooltip" title="Reply">
                <span class="Icon  Icon--reply"></span>
                <span class="u-hiddenVisually">Reply</span>
              </div>
            </button>
          </div>
          <div class="ProfileTweet-action ProfileTweet-action--retweet js-toggleState js-toggleRt">
            <button class="ProfileTweet-actionButton  js-actionButton js-actionRetweet" data-modal="ProfileTweet-retweet" type="button">
              <div class="IconContainer js-tooltip" title="Retweet">
                <span class="Icon  Icon--retweet"></span>
                <span class="u-hiddenVisually">Retweet</span>
              </div>
                <div class="IconTextContainer">
                  <span class="ProfileTweet-actionCount ProfileTweet-actionCount--isZero">
                    <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true"></span>
                  </span>
                </div>
            </button><button class="ProfileTweet-actionButtonUndo js-actionButton js-actionRetweet" data-modal="ProfileTweet-retweet" type="button">
              <div class="IconContainer js-tooltip" title="Undo retweet">
                <span class="Icon  Icon--retweet"></span>
                <span class="u-hiddenVisually">Retweeted</span>
              </div>
                <div class="IconTextContainer">
                  <span class="ProfileTweet-actionCount ProfileTweet-actionCount--isZero">
                    <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true"></span>
                  </span>
                </div>
            </button>
          </div>
          <div class="ProfileTweet-action ProfileTweet-action--favorite js-toggleState ">
            <button class="ProfileTweet-actionButton js-actionButton js-actionFavorite" type="button">
              <div class="IconContainer js-tooltip" title="Favorite">
                <span class="Icon  Icon--favorite"></span>
                <span class="u-hiddenVisually">Favorite</span>
              </div>
                <div class="IconTextContainer">
                  <span class="ProfileTweet-actionCount ProfileTweet-actionCount--isZero">
                      <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true"></span>
                  </span>
                </div>
            </button><button class="ProfileTweet-actionButtonUndo u-linkClean js-actionButton js-actionFavorite" type="button">
              <div class="IconContainer js-tooltip" title="Undo favorite">
                <span class="Icon  Icon--favorite"></span>
                  <span class="u-hiddenVisually">Favorited</span>
              </div>
                <div class="IconTextContainer">
                  <span class="ProfileTweet-actionCount ProfileTweet-actionCount--isZero">
                      <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true"></span>
                  </span>
                </div>
            </button>
          </div>
            <div class="ProfileTweet-action ProfileTweet-action--more js-more-ProfileTweet-actions">
              <div class="dropdown">
      <button aria-haspopup="true" class="ProfileTweet-actionButton u-textUserColorHover dropdown-toggle js-dropdown-toggle" type="button">
          <div class="IconContainer js-tooltip" title="More">
            <span class="Icon  Icon--dots"></span>
            <span class="u-hiddenVisually">More</span>
          </div>
      </button>
      <div class="dropdown-menu">
      <div class="dropdown-caret">
        <div class="caret-outer"></div>
        <div class="caret-inner"></div>
      </div>
      <ul>
          <li class="share-via-dm js-actionShareViaDM" data-nav="share_tweet_dm">
            <button type="button" class="dropdown-link">Share via Direct Message</button>
          </li>

          <li class="copy-link-to-tweet js-actionCopyLinkToTweet">
            <button type="button" class="dropdown-link">Copy link to Tweet</button>
          </li>
            <li class="embed-link js-actionEmbedTweet" data-nav="embed_tweet">
              <button type="button" class="dropdown-link">Embed Tweet</button>
            </li>
              <li class="mute-user-item pretty-link"><button type="button" class="dropdown-link">Mute</button></li>
      <li class="unmute-user-item pretty-link"><button type="button" class="dropdown-link">Unmute</button></li>
            <li class="block-link js-actionBlock" data-nav="block">
              <button type="button" class="dropdown-link">Block</button>
            </li>
            <li class="unblock-link js-actionUnblock" data-nav="unblock">
              <button type="button" class="dropdown-link">Unblock</button>
            </li>
            <li class="report-link js-actionReport" data-nav="report">
              <button type="button" class="dropdown-link">
                Report
              </button>
            </li>
      </ul>
    </div>
    </div>
            </div>
        </div>
    </div>
        </div>
      </div>
    </li>
`
	streamRawHTML = `        <ol class="stream-items js-navigable-stream" id="stream-items-id">
            <li class="js-stream-item stream-item stream-item expanding-stream-item js-pinned
            " data-item-id="637274908041605120" id="stream-item-tweet-637274908041605120" data-item-type="tweet">
            <ol role="presentation" class="expanded-conversation expansion-container js-expansion-container js-navigable-stream">
            <li role="presentation" class="original-tweet-container">
            <div class="tweet original-tweet js-original-tweet js-stream-tweet js-actionable-tweet js-profile-popup-actionable
            my-tweet
            has-cards
            cards-forward
            user-pinned
            with-non-tweet-action-follow-button
            " data-tweet-id="637274908041605120" data-disclosure-type="" data-item-id="637274908041605120" data-permalink-path="/onetruecathal/status/637274908041605120" data-screen-name="onetruecathal" data-name="Cathal Garvey" data-user-id="16066596" data-has-cards="true" data-card2-type="summary" data-expanded-footer="<div class=&quot;js-tweet-details-fixer tweet-details-fixer&quot;>
            <div class=&quot;entities-media-container js-media-container&quot; style=&quot;min-height:0px&quot;>
            </div>
            <div class=&quot;js-machine-translated-tweet-container&quot;></div>
            <div class=&quot;js-tweet-stats-container tweet-stats-container &quot;>
            </div>
            <div class=&quot;client-and-actions&quot;>
            <span class=&quot;metadata&quot;>
            <span>3:45 PM - 28 Aug 2015</span>
            &amp;middot; <a class=&quot;permalink-link js-permalink js-nav&quot; href=&quot;/onetruecathal/status/637274908041605120&quot;  tabindex=&quot;-1&quot;>Details</a>
            </span>
            </div>
            </div>
            " data-mentions="FormaBiolabs" data-you-follow="false" data-follows-you="false" data-you-block="false" data-tfb-view="/i/tfb/v1/quick_promote/637274908041605120">
            <div class="context">
            <div class="tweet-context with-icn
            pinned">
            <span class="Icon Icon--small Icon--pinned u-textUserColor"></span>
            <span class="js-pinned-text" data-aria-label-part="">Pinned Tweet</span>
            </div>
            </div>
            <div class="content">
            <div class="stream-item-header">
            <a class="account-group js-account-group js-action-profile js-user-profile-link js-nav" href="/onetruecathal" data-user-id="16066596">
            <img class="avatar js-action-profile-avatar" src="https://pbs.twimg.com/profile_images/574863646813106176/27q-jhEb_bigger.png" alt="">
            <strong class="fullname js-action-profile-name show-popup-with-id" data-aria-label-part="">Cathal Garvey</strong>
            <span>‏</span><span class="username js-action-profile-name" data-aria-label-part=""><s>@</s><b>onetruecathal</b></span>
            </a>
            <small class="time">
            <a data-original-title="3:45 PM - 28 Aug 2015" href="/onetruecathal/status/637274908041605120" class="tweet-timestamp js-permalink js-nav js-tooltip"><span class="_timestamp js-short-timestamp " data-aria-label-part="last" data-time="1440773152" data-time-ms="1440773152000" data-long-form="true">Aug 28</span></a>
            </small>
            <span class="Tweet-geo u-floatRight js-tooltip" title="Cork, Ireland">
            <a class="ProfileTweet-actionButton u-linkClean js-nav js-geo-pivot-link" href="/search?q=place%3Ac500c6f8494a90ac" role="button" data-place-id="c500c6f8494a90ac">
            <span class="Icon Icon--geo"></span>
            <span class="u-hiddenVisually">Cork, Ireland</span>
            </a>
            </span>
            </div>
            <p class="TweetTextSize TweetTextSize--26px js-tweet-text tweet-text" data-aria-label-part="0" lang="en">Join us Thursday next at 6pm for <a href="/FormaBiolabs" class="twitter-atreply pretty-link js-nav" dir="ltr"><s>@</s><b>FormaBiolabs</b></a>' official (but relaxed!) opening party: <a href="http://t.co/lhrrPxz5Fy" rel="nofollow" dir="ltr" data-expanded-url="http://www.eventbrite.ie/e/forma-labs-opening-ceremony-tickets-18264954972" class="twitter-timeline-link" target="_blank" title="http://www.eventbrite.ie/e/forma-labs-opening-ceremony-tickets-18264954972"><span class="tco-ellipsis"></span><span class="invisible">http://www.</span><span class="js-display-url">eventbrite.ie/e/forma-labs-o</span><span class="invisible">pening-ceremony-tickets-18264954972</span><span class="tco-ellipsis"><span class="invisible">&nbsp;</span>…</span></a> <a href="/hashtag/science?src=hash" data-query-source="hashtag_click" class="twitter-hashtag pretty-link js-nav" dir="ltr"><s>#</s><b>science</b></a> <a href="/hashtag/cork?src=hash" data-query-source="hashtag_click" class="twitter-hashtag pretty-link js-nav" dir="ltr"><s>#</s><b>cork</b></a> <a href="/hashtag/scicomm?src=hash" data-query-source="hashtag_click" class="twitter-hashtag pretty-link js-nav" dir="ltr"><s>#</s><b>scicomm</b></a></p>
            <div class="card2 js-media-container portrait" data-card2-name="summary">
            <div data-watched="true" class="js-macaw-cards-iframe-container" data-src="/i/cards/tfw/v1/637274908041605120?cardname=summary&amp;earned=true" data-autoplay-src="/i/cards/tfw/v1/637274908041605120?cardname=summary&amp;earned=true" data-card-name="summary" data-card-url="http://t.co/lhrrPxz5Fy" data-publisher-id="5625972" data-creator-id="" data-amplify-content-id="" data-amplify-playlist-url="" data-full-card-iframe-url="/i/cards/tfw/v1/637274908041605120?cardname=summary&amp;earned=true" data-has-autoplayable-media="false">
            <iframe allowfullscreen="" src="https://twitter.com/i/cards/tfw/v1/637274908041605120?cardname=summary&amp;earned=true#xdm_e=https%3A%2F%2Ftwitter.com&amp;xdm_c=default3756&amp;xdm_p=1" scrolling="no" style="display: block; margin: 0px; padding: 0px; border: 0px none;" id="xdm_default3756_provider" frameborder="0" height="129" width="100%"></iframe></div>
            </div>
            <div class="expanded-content js-tweet-details-dropdown">
            <div class="js-tweet-details-fixer tweet-details-fixer">
            <div class="entities-media-container js-media-container" style="min-height:0px">
            </div>
            <div class="js-machine-translated-tweet-container"></div>
            <div class="js-tweet-stats-container tweet-stats-container ">
            </div>
            <div class="client-and-actions">
            <span class="metadata">
            <span>3:45 PM - 28 Aug 2015</span>
            · <a class="permalink-link js-permalink js-nav" href="/onetruecathal/status/637274908041605120" tabindex="-1">Details</a>
            </span>
            </div>
            </div>
            </div>
            <div class="stream-item-footer">
            <span class="ProfileTweet-action--reply u-hiddenVisually"></span>
            <span class="ProfileTweet-action--retweet u-hiddenVisually">
            <span class="ProfileTweet-actionCount" data-tweet-stat-count="12">
            <span class="ProfileTweet-actionCountForAria" data-aria-label-part="">12 retweets</span>
            </span>
            </span>
            <span class="ProfileTweet-action--favorite u-hiddenVisually">
            <span class="ProfileTweet-actionCount" data-tweet-stat-count="9">
            <span class="ProfileTweet-actionCountForAria" data-aria-label-part="">9 favorites</span>
            </span>
            </span>
            <div role="group" aria-label="Tweet actions" class="ProfileTweet-actionList u-cf js-actions ">
            <div class="ProfileTweet-action ProfileTweet-action--reply">
            <button class="ProfileTweet-actionButton u-textUserColorHover js-actionButton js-actionReply" data-modal="ProfileTweet-reply" type="button">
            <div class="IconContainer js-tooltip" title="Reply">
            <span class="Icon  Icon--reply"></span>
            <span class="u-hiddenVisually">Reply</span>
            </div>
            </button>
            </div>
            <div class="ProfileTweet-action ProfileTweet-action--retweet js-toggleState js-toggleRt">
            <button class="ProfileTweet-actionButton  is-disabled js-disableTweetAction js-actionButton js-actionRetweet" disabled="" data-modal="ProfileTweet-retweet" type="button">
            <div class="IconContainer js-tooltip">
            <span class="Icon  Icon--retweet"></span>
            <span class="u-hiddenVisually">Retweet</span>
            </div>
            <div class="IconTextContainer">
            <span class="ProfileTweet-actionCount">
            <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true">12</span>
            </span>
            </div>
            </button><button class="ProfileTweet-actionButtonUndo js-actionButton js-actionRetweet" data-modal="ProfileTweet-retweet" type="button">
            <div class="IconContainer js-tooltip" title="Undo retweet">
            <span class="Icon  Icon--retweet"></span>
            <span class="u-hiddenVisually">Retweeted</span>
            </div>
            <div class="IconTextContainer">
            <span class="ProfileTweet-actionCount">
            <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true">12</span>
            </span>
            </div>
            </button>
            </div>
            <div class="ProfileTweet-action ProfileTweet-action--favorite js-toggleState ">
            <button class="ProfileTweet-actionButton js-actionButton js-actionFavorite" type="button">
            <div class="IconContainer js-tooltip" title="Favorite">
            <span class="Icon  Icon--favorite"></span>
            <span class="u-hiddenVisually">Favorite</span>
            </div>
            <div class="IconTextContainer">
            <span class="ProfileTweet-actionCount">
            <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true">9</span>
            </span>
            </div>
            </button><button class="ProfileTweet-actionButtonUndo u-linkClean js-actionButton js-actionFavorite" type="button">
            <div class="IconContainer js-tooltip" title="Undo favorite">
            <span class="Icon  Icon--favorite"></span>
            <span class="u-hiddenVisually">Favorited</span>
            </div>
            <div class="IconTextContainer">
            <span class="ProfileTweet-actionCount">
            <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true">9</span>
            </span>
            </div>
            </button>
            </div>
            <div class="ProfileTweet-action ProfileTweet-action--analytics">
            <button class="ProfileTweet-actionButton u-textUserColorHover js-actionButton js-actionQuickPromote" type="button">
            <div class="IconContainer js-tooltip" title="View Tweet activity">
            <span class="Icon  Icon--analytics"></span>
            <span class="u-hiddenVisually">View Tweet activity</span>
            </div>
            </button>
            </div>
            <div class="ProfileTweet-action ProfileTweet-action--more js-more-ProfileTweet-actions">
            <div class="dropdown">
            <button aria-haspopup="true" class="ProfileTweet-actionButton u-textUserColorHover dropdown-toggle js-dropdown-toggle" type="button">
            <div class="IconContainer js-tooltip" title="More">
            <span class="Icon  Icon--dots"></span>
            <span class="u-hiddenVisually">More</span>
            </div>
            </button>
            <div class="dropdown-menu">
            <div class="dropdown-caret">
            <div class="caret-outer"></div>
            <div class="caret-inner"></div>
            </div>
            <ul>
            <li class="copy-link-to-tweet js-actionCopyLinkToTweet">
            <button type="button" class="dropdown-link">Copy link to Tweet</button>
            </li>
            <li class="embed-link js-actionEmbedTweet" data-nav="embed_tweet">
            <button type="button" class="dropdown-link">Embed Tweet</button>
            </li>
            <li class="user-pin-tweet js-actionPinTweet" data-nav="user_pin_tweet">
            <button type="button" class="dropdown-link">Pin to your profile page</button>
            </li>
            <li class="user-unpin-tweet js-actionPinTweet" data-nav="user_unpin_tweet">
            <button type="button" class="dropdown-link">Unpin from profile page</button>
            </li>
            <li class="js-actionDelete">
            <button type="button" class="dropdown-link">Delete Tweet</button>
            </li>
            </ul>
            </div>
            </div>
            </div>
            </div>
            </div>
            </div>
            </div>
            </li>
            </ol>
            </li>
            <li class="js-stream-item stream-item stream-item expanding-stream-item
            " data-item-id="642042919466180608" id="stream-item-tweet-642042919466180608" data-item-type="tweet">
            <div class="tweet original-tweet js-original-tweet js-stream-tweet js-actionable-tweet js-profile-popup-actionable
            retweeted
            with-non-tweet-action-follow-button
            " data-tweet-id="642042919466180608" data-disclosure-type="" data-item-id="642042919466180608" data-permalink-path="/csoghoian/status/642042919466180608" data-my-retweet-id="642046037167853569" data-retweet-id="642046037167853569" data-screen-name="csoghoian" data-name="Christopher Soghoian" data-user-id="14669471" data-expanded-footer="<div class=&quot;js-tweet-details-fixer tweet-details-fixer&quot;>
            <div class=&quot;entities-media-container js-media-container&quot; style=&quot;min-height:0px&quot;>
            </div>
            <div class=&quot;js-machine-translated-tweet-container&quot;></div>
            <div class=&quot;js-tweet-stats-container tweet-stats-container &quot;>
            </div>
            <div class=&quot;client-and-actions&quot;>
            <span class=&quot;metadata&quot;>
            <span>7:32 PM - 10 Sep 2015</span>
            &amp;middot; <a class=&quot;permalink-link js-permalink js-nav&quot; href=&quot;/csoghoian/status/642042919466180608&quot;  tabindex=&quot;-1&quot;>Details</a>
            </span>
            </div>
            </div>
            " data-retweeter="onetruecathal" data-you-follow="true" data-follows-you="false" data-you-block="false">
            <div class="context">
            <div class="tweet-context with-icn
            ">
            <span class="Icon Icon--small Icon--retweeted"></span>
            <span class="js-retweet-text"><a class="pretty-link js-user-profile-link" href="/onetruecathal" data-user-id="16066596"><b>Cathal Garvey</b></a> retweeted</span>
            </div>
            </div>
            <div class="content">
            <div class="stream-item-header">
            <a class="account-group js-account-group js-action-profile js-user-profile-link js-nav" href="/csoghoian" data-user-id="14669471">
            <img class="avatar js-action-profile-avatar" src="https://pbs.twimg.com/profile_images/378800000723950844/386b29d0ac8506a0d916f7fa6634e4c8_bigger.jpeg" alt="">
            <strong class="fullname js-action-profile-name show-popup-with-id" data-aria-label-part="">Christopher Soghoian</strong>
            <span>‏</span><span class="username js-action-profile-name" data-aria-label-part=""><s>@</s><b>csoghoian</b></span>
            </a>
            <small class="time">
            <a href="/csoghoian/status/642042919466180608" class="tweet-timestamp js-permalink js-nav js-tooltip" title="7:32 PM - 10 Sep 2015"><span class="_timestamp js-short-timestamp js-relative-timestamp" data-time="1441909935" data-time-ms="1441909935000" data-long-form="true" aria-hidden="true">3h</span><span class="u-hiddenVisually" data-aria-label-part="last">3 hours ago</span></a>
            </small>
            </div>
            <p class="TweetTextSize TweetTextSize--16px js-tweet-text tweet-text" data-aria-label-part="0" lang="en">Shorter FBI Director: I love strong encryption. I want strong encryption to protect my data. I just don't want you to have strong encryption</p>
            <div class="expanded-content js-tweet-details-dropdown">
            </div>
            <div class="stream-item-footer">
            <span class="ProfileTweet-action--reply u-hiddenVisually"></span>
            <span class="ProfileTweet-action--retweet u-hiddenVisually">
            <span class="ProfileTweet-actionCount" data-tweet-stat-count="40">
            <span class="ProfileTweet-actionCountForAria" data-aria-label-part="">40 retweets</span>
            </span>
            </span>
            <span class="ProfileTweet-action--favorite u-hiddenVisually">
            <span class="ProfileTweet-actionCount" data-tweet-stat-count="39">
            <span class="ProfileTweet-actionCountForAria" data-aria-label-part="">39 favorites</span>
            </span>
            </span>
            <div role="group" aria-label="Tweet actions" class="ProfileTweet-actionList u-cf js-actions ">
            <div class="ProfileTweet-action ProfileTweet-action--reply">
            <button class="ProfileTweet-actionButton u-textUserColorHover js-actionButton js-actionReply" data-modal="ProfileTweet-reply" type="button">
            <div class="IconContainer js-tooltip" title="Reply">
            <span class="Icon  Icon--reply"></span>
            <span class="u-hiddenVisually">Reply</span>
            </div>
            </button>
            </div>
            <div class="ProfileTweet-action ProfileTweet-action--retweet js-toggleState js-toggleRt">
            <button class="ProfileTweet-actionButton  js-actionButton js-actionRetweet" data-modal="ProfileTweet-retweet" type="button">
            <div class="IconContainer js-tooltip" title="Retweet">
            <span class="Icon  Icon--retweet"></span>
            <span class="u-hiddenVisually">Retweet</span>
            </div>
            <div class="IconTextContainer">
            <span class="ProfileTweet-actionCount">
            <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true">40</span>
            </span>
            </div>
            </button><button class="ProfileTweet-actionButtonUndo js-actionButton js-actionRetweet" data-modal="ProfileTweet-retweet" type="button">
            <div class="IconContainer js-tooltip" title="Undo retweet">
            <span class="Icon  Icon--retweet"></span>
            <span class="u-hiddenVisually">Retweeted</span>
            </div>
            <div class="IconTextContainer">
            <span class="ProfileTweet-actionCount">
            <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true">40</span>
            </span>
            </div>
            </button>
            </div>
            <div class="ProfileTweet-action ProfileTweet-action--favorite js-toggleState ">
            <button class="ProfileTweet-actionButton js-actionButton js-actionFavorite" type="button">
            <div class="IconContainer js-tooltip" title="Favorite">
            <span class="Icon  Icon--favorite"></span>
            <span class="u-hiddenVisually">Favorite</span>
            </div>
            <div class="IconTextContainer">
            <span class="ProfileTweet-actionCount">
            <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true">39</span>
            </span>
            </div>
            </button><button class="ProfileTweet-actionButtonUndo u-linkClean js-actionButton js-actionFavorite" type="button">
            <div class="IconContainer js-tooltip" title="Undo favorite">
            <span class="Icon  Icon--favorite"></span>
            <span class="u-hiddenVisually">Favorited</span>
            </div>
            <div class="IconTextContainer">
            <span class="ProfileTweet-actionCount">
            <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true">39</span>
            </span>
            </div>
            </button>
            </div>
            <div class="ProfileTweet-action ProfileTweet-action--more js-more-ProfileTweet-actions">
            <div class="dropdown">
            <button aria-haspopup="true" class="ProfileTweet-actionButton u-textUserColorHover dropdown-toggle js-dropdown-toggle" type="button">
            <div class="IconContainer js-tooltip" title="More">
            <span class="Icon  Icon--dots"></span>
            <span class="u-hiddenVisually">More</span>
            </div>
            </button>
            <div class="dropdown-menu">
            <div class="dropdown-caret">
            <div class="caret-outer"></div>
            <div class="caret-inner"></div>
            </div>
            <ul>
            <li class="share-via-dm js-actionShareViaDM" data-nav="share_tweet_dm">
            <button type="button" class="dropdown-link">Share via Direct Message</button>
            </li>
            <li class="copy-link-to-tweet js-actionCopyLinkToTweet">
            <button type="button" class="dropdown-link">Copy link to Tweet</button>
            </li>
            <li class="embed-link js-actionEmbedTweet" data-nav="embed_tweet">
            <button type="button" class="dropdown-link">Embed Tweet</button>
            </li>
            <li class="mute-user-item pretty-link"><button type="button" class="dropdown-link">Mute</button></li>
            <li class="unmute-user-item pretty-link"><button type="button" class="dropdown-link">Unmute</button></li>
            <li class="block-link js-actionBlock" data-nav="block">
            <button type="button" class="dropdown-link">Block</button>
            </li>
            <li class="unblock-link js-actionUnblock" data-nav="unblock">
            <button type="button" class="dropdown-link">Unblock</button>
            </li>
            <li class="report-link js-actionReport" data-nav="report">
            <button type="button" class="dropdown-link">
            Report
            </button>
            </li>
            </ul>
            </div>
            </div>
            </div>
            </div>
            </div>
            </div>
            </div>
            </li>
            <li class="js-stream-item stream-item stream-item expanding-stream-item
            " data-item-id="642045936995336193" id="stream-item-tweet-642045936995336193" data-item-type="tweet">
            <div class="tweet original-tweet js-original-tweet js-stream-tweet js-actionable-tweet js-profile-popup-actionable
            my-tweet
            with-non-tweet-action-follow-button
            " data-tweet-id="642045936995336193" data-disclosure-type="" data-item-id="642045936995336193" data-permalink-path="/onetruecathal/status/642045936995336193" data-screen-name="onetruecathal" data-name="Cathal Garvey" data-user-id="16066596" data-expanded-footer="<div class=&quot;js-tweet-details-fixer tweet-details-fixer&quot;>
            <div class=&quot;js-machine-translated-tweet-container&quot;></div>
            <div class=&quot;js-tweet-stats-container tweet-stats-container &quot;>
            </div>
            <div class=&quot;client-and-actions&quot;>
            <span class=&quot;metadata&quot;>
            <span>7:44 PM - 10 Sep 2015</span>
            &amp;middot; <a class=&quot;permalink-link js-permalink js-nav&quot; href=&quot;/onetruecathal/status/642045936995336193&quot;  tabindex=&quot;-1&quot;>Details</a>
            </span>
            </div>
            </div>
            " data-you-follow="false" data-follows-you="false" data-you-block="false" data-tfb-view="/i/tfb/v1/quick_promote/642045936995336193">
            <div class="context">
            </div>
            <div class="content">
            <div class="stream-item-header">
            <a class="account-group js-account-group js-action-profile js-user-profile-link js-nav" href="/onetruecathal" data-user-id="16066596">
            <img class="avatar js-action-profile-avatar" src="https://pbs.twimg.com/profile_images/574863646813106176/27q-jhEb_bigger.png" alt="">
            <strong class="fullname js-action-profile-name show-popup-with-id">Cathal Garvey</strong>
            <span>‏</span><span class="username js-action-profile-name"><s>@</s><b>onetruecathal</b></span>
            </a>
            <small class="time">
            <a data-original-title="7:44 PM - 10 Sep 2015" href="/onetruecathal/status/642045936995336193" class="tweet-timestamp js-permalink js-nav js-tooltip"><span class="_timestamp js-short-timestamp js-relative-timestamp" data-time="1441910654" data-time-ms="1441910654000" data-long-form="true" aria-hidden="true">3h</span><span class="u-hiddenVisually" data-aria-label-part="last">3 hours ago</span></a>
            </small>
            </div>
            <p class="u-hiddenVisually" aria-hidden="true" data-aria-label-part="1">Cathal Garvey retweeted Frank N. Foode</p>
            <p class="TweetTextSize TweetTextSize--16px js-tweet-text tweet-text" data-aria-label-part="4" lang="en">Awesome news for <a href="/hashtag/celiac?src=hash" data-query-source="hashtag_click" class="twitter-hashtag pretty-link js-nav" dir="ltr"><s>#</s><b>celiac</b></a> peeps: soon you could be enjoying 4-5 slices of decent bread! <a href="https://t.co/N6jJIOrgNG" rel="nofollow" dir="ltr" data-expanded-url="https://twitter.com/franknfoode/status/642042436240429057" class="twitter-timeline-link u-hidden" target="_blank" title="https://twitter.com/franknfoode/status/642042436240429057"><span class="tco-ellipsis"></span><span class="invisible">https://</span><span class="js-display-url">twitter.com/franknfoode/st</span><span class="invisible">atus/642042436240429057</span><span class="tco-ellipsis"><span class="invisible">&nbsp;</span>…</span></a></p>
            <p class="u-hiddenVisually" aria-hidden="true" data-aria-label-part="3">Cathal Garvey added,</p>
            <div class="QuoteTweet
            u-block js-tweet-details-fixer">
            <div class="QuoteTweet-container">
            <a class="QuoteTweet-link" href="/franknfoode/status/638443323078918144" aria-hidden="true">
            </a>
            <div class="QuoteTweet-innerContainer u-cf js-permalink js-media-container" data-item-id="638443323078918144" data-item-type="tweet" data-screen-name="franknfoode" data-user-id="74625926" href="/franknfoode/status/638443323078918144" tabindex="0">
            <div class="tweet-content">
            <div class="QuoteTweet-mediaContainer  with-text">
            <div class="multi-photos photos-1">
            <div class="multi-photo photo-1">
            <img src="https://pbs.twimg.com/media/CNw0H3nWgAApYb_.jpg" style="left: -42%; height: 100%; width: auto;">
            </div>
            </div>
            </div>
            <div class="QuoteTweet-authorAndText u-alignTop">
            <span class="QuoteTweet-originalAuthor u-cf u-textTruncate stream-item-header js-user-profile-link">
            <b class="QuoteTweet-fullname u-linkComplex-target">Frank N. Foode</b>
            <span class="QuoteTweet-screenname u-dir" dir="ltr">
            <span class="at">@</span>franknfoode
            </span>
            </span>
            <div class="QuoteTweet-text tweet-text u-dir" data-aria-label-part="2" dir="ltr" lang="en">Guess which bread was made with <a data-query-source="hashtag_click" class="twitter-hashtag pretty-link js-nav" dir="ltr"><s>#</s><b>glutenfree</b></a> <a data-query-source="hashtag_click" class="twitter-hashtag pretty-link js-nav" dir="ltr"><s>#</s><b>GMO</b></a> wheat? Answer here: <a rel="nofollow" dir="ltr" data-expanded-url="http://www.biofortified.org/2015/08/gluten-free-gm-wheat-can-help-celiac-patients/" class="twitter-timeline-link" target="_blank" title="http://www.biofortified.org/2015/08/gluten-free-gm-wheat-can-help-celiac-patients/"><span class="tco-ellipsis"></span><span class="invisible">http://www.</span><span class="js-display-url">biofortified.org/2015/08/gluten</span><span class="invisible">-free-gm-wheat-can-help-celiac-patients/</span><span class="tco-ellipsis"><span class="invisible">&nbsp;</span>…</span></a> via <a class="twitter-atreply pretty-link js-nav" dir="ltr"><s>@</s><b>dyescience</b></a> <a class="twitter-timeline-link u-hidden" data-pre-embedded="true" dir="ltr">pic.twitter.com/b8Fvb06hPa</a></div>
            </div>
            </div>
            </div>
            </div>
            </div>
            <div class="expanded-content js-tweet-details-dropdown">
            </div>
            <div class="stream-item-footer">
            <span class="ProfileTweet-action--reply u-hiddenVisually"></span>
            <span class="ProfileTweet-action--retweet u-hiddenVisually">
            <span class="ProfileTweet-actionCount" aria-hidden="true" data-tweet-stat-count="0">
            <span class="ProfileTweet-actionCountForAria">0 retweets</span>
            </span>
            </span>
            <span class="ProfileTweet-action--favorite u-hiddenVisually">
            <span class="ProfileTweet-actionCount" aria-hidden="true" data-tweet-stat-count="0">
            <span class="ProfileTweet-actionCountForAria">0 favorites</span>
            </span>
            </span>
            <div role="group" aria-label="Tweet actions" class="ProfileTweet-actionList u-cf js-actions ">
            <div class="ProfileTweet-action ProfileTweet-action--reply">
            <button class="ProfileTweet-actionButton u-textUserColorHover js-actionButton js-actionReply" data-modal="ProfileTweet-reply" type="button">
            <div class="IconContainer js-tooltip" title="Reply">
            <span class="Icon  Icon--reply"></span>
            <span class="u-hiddenVisually">Reply</span>
            </div>
            </button>
            </div>
            <div class="ProfileTweet-action ProfileTweet-action--retweet js-toggleState js-toggleRt">
            <button class="ProfileTweet-actionButton  is-disabled js-disableTweetAction js-actionButton js-actionRetweet" disabled="" data-modal="ProfileTweet-retweet" type="button">
            <div class="IconContainer js-tooltip">
            <span class="Icon  Icon--retweet"></span>
            <span class="u-hiddenVisually">Retweet</span>
            </div>
            <div class="IconTextContainer">
            <span class="ProfileTweet-actionCount ProfileTweet-actionCount--isZero">
            <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true"></span>
            </span>
            </div>
            </button><button class="ProfileTweet-actionButtonUndo js-actionButton js-actionRetweet" data-modal="ProfileTweet-retweet" type="button">
            <div class="IconContainer js-tooltip" title="Undo retweet">
            <span class="Icon  Icon--retweet"></span>
            <span class="u-hiddenVisually">Retweeted</span>
            </div>
            <div class="IconTextContainer">
            <span class="ProfileTweet-actionCount ProfileTweet-actionCount--isZero">
            <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true"></span>
            </span>
            </div>
            </button>
            </div>
            <div class="ProfileTweet-action ProfileTweet-action--favorite js-toggleState ">
            <button class="ProfileTweet-actionButton js-actionButton js-actionFavorite" type="button">
            <div class="IconContainer js-tooltip" title="Favorite">
            <span class="Icon  Icon--favorite"></span>
            <span class="u-hiddenVisually">Favorite</span>
            </div>
            <div class="IconTextContainer">
            <span class="ProfileTweet-actionCount ProfileTweet-actionCount--isZero">
            <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true"></span>
            </span>
            </div>
            </button><button class="ProfileTweet-actionButtonUndo u-linkClean js-actionButton js-actionFavorite" type="button">
            <div class="IconContainer js-tooltip" title="Undo favorite">
            <span class="Icon  Icon--favorite"></span>
            <span class="u-hiddenVisually">Favorited</span>
            </div>
            <div class="IconTextContainer">
            <span class="ProfileTweet-actionCount ProfileTweet-actionCount--isZero">
            <span class="ProfileTweet-actionCountForPresentation" aria-hidden="true"></span>
            </span>
            </div>
            </button>
            </div>
            <div class="ProfileTweet-action ProfileTweet-action--analytics">
            <button class="ProfileTweet-actionButton u-textUserColorHover js-actionButton js-actionQuickPromote" type="button">
            <div class="IconContainer js-tooltip" title="View Tweet activity">
            <span class="Icon  Icon--analytics"></span>
            <span class="u-hiddenVisually">View Tweet activity</span>
            </div>
            </button>
            </div>
            <div class="ProfileTweet-action ProfileTweet-action--more js-more-ProfileTweet-actions">
            <div class="dropdown">
            <button aria-haspopup="true" class="ProfileTweet-actionButton u-textUserColorHover dropdown-toggle js-dropdown-toggle" type="button">
            <div class="IconContainer js-tooltip" title="More">
            <span class="Icon  Icon--dots"></span>
            <span class="u-hiddenVisually">More</span>
            </div>
            </button>
            <div class="dropdown-menu">
            <div class="dropdown-caret">
            <div class="caret-outer"></div>
            <div class="caret-inner"></div>
            </div>
            <ul>
            <li class="copy-link-to-tweet js-actionCopyLinkToTweet">
            <button type="button" class="dropdown-link">Copy link to Tweet</button>
            </li>
            <li class="embed-link js-actionEmbedTweet" data-nav="embed_tweet">
            <button type="button" class="dropdown-link">Embed Tweet</button>
            </li>
            <li class="user-pin-tweet js-actionPinTweet" data-nav="user_pin_tweet">
            <button type="button" class="dropdown-link">Pin to your profile page</button>
            </li>
            <li class="user-unpin-tweet js-actionPinTweet" data-nav="user_unpin_tweet">
            <button type="button" class="dropdown-link">Unpin from profile page</button>
            </li>
            <li class="js-actionDelete">
            <button type="button" class="dropdown-link">Delete Tweet</button>
            </li>
            </ul>
            </div>
            </div>
            </div>
            </div>
            </div>
            </div>
            </div>
            </li>
            </ol>`
)
