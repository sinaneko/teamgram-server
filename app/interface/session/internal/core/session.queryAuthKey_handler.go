// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/session/session"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
)

// SessionQueryAuthKey
// session.queryAuthKey auth_key_id:long = AuthKeyInfo;
func (c *SessionCore) SessionQueryAuthKey(in *session.TLSessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	key, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionQueryAuthKey(c.ctx, &authsession.TLAuthsessionQueryAuthKey{
		AuthKeyId: in.GetAuthKeyId(),
	})
	if err != nil {
		c.Logger.Errorf("session.queryAuthKey - error: %v", err)
		return nil, err
	}

	return key, nil
}
