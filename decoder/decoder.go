/*
This package contains the decoder definitions that are used to
decode the messages from the protocol supported by the decoder
and convert them to the CommonEntity format.

All decoders should implement the Decoder interface.

See HttpDecoder for a very simple example.
*/

package decoder

import "github.com/csizsek/prism/entity"

/*
Decoder is the interface that used to define the Decode method, this
interface should be implemented for each protocol that is used to
feed data into Prism.
*/
type Decoder interface {
	Decode(*entity.Entity) *entity.CommonEntity
}
