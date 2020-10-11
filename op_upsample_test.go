package gorgonia

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorgonia.org/tensor"
)

type testData struct {
	insh    []int
	indata  []float64
	outsh   []int
	outdata []float64
}

var data testData = testData{
	insh:    []int{1, 5, 5, 8},
	indata:  []float64{3.588184118270874, 6.765244483947754, 8.703499794006348, 7.832754611968994, 7.205557346343994, 2.2351908683776855, 9.140625, 1.8968262672424316, 7.579838752746582, 2.0422589778900146, 1.7213886976242065, 8.008468627929688, 1.8242313861846924, 0.15505588054656982, 7.218662738800049, 6.959044456481934, 5.268470287322998, 3.0470921993255615, 6.904945373535156, 0.048324376344680786, 5.862209796905518, 9.752281188964844, 9.729859352111816, 6.917296409606934, 5.983068466186523, 8.5852689743042, 8.877313613891602, 0.06957779079675674, 8.360713958740234, 1.0679903030395508, 1.2100963592529297, 9.838610649108887, 2.814711809158325, 6.38685941696167, 8.260359764099121, 2.532329559326172, 7.589834213256836, 1.294305443763733, 0.6550663113594055, 0.8506288528442383, 2.1071043014526367, 8.425798416137695, 8.408304214477539, 7.394162178039551, 2.590216875076294, 0.7849066257476807, 2.676325559616089, 8.072299003601074, 3.4864327907562256, 2.422727346420288, 9.429472923278809, 2.553463935852051, 5.96559476852417, 7.482579231262207, 2.3580377101898193, 1.5595018863677979, 0.034589122980833054, 3.0223019123077393, 4.7403154373168945, 0.2549790143966675, 6.765586853027344, 6.91905403137207, 1.7624300718307495, 1.8117458820343018, 5.739025592803955, 7.67076301574707, 8.772150039672852, 2.1648266315460205, 2.8233001232147217, 3.455390214920044, 1.954986333847046, 6.2339935302734375, 8.180193901062012, 3.4818568229675293, 0.5501848459243774, 7.886703014373779, 3.614240884780884, 7.023927211761475, 0.4005505442619324, 2.0974090099334717, 6.166986465454102, 4.413336753845215, 7.7545647621154785, 5.656180381774902, 8.241524696350098, 0.1758895069360733, 7.527948379516602, 2.4103832244873047, 2.5865209102630615, 8.264690399169922, 4.0821614265441895, 4.711780548095703, 8.985666275024414, 3.614298105239868, 2.9979336261749268, 1.0938977003097534, 0.6748584508895874, 8.776729583740234, 7.792627811431885, 4.120683670043945, 8.067517280578613, 6.481339931488037, 3.7537691593170166, 1.5674660205841064, 4.622176647186279, 6.572947025299072, 6.97075891494751, 0.5755885243415833, 9.231983184814453, 5.7552490234375, 6.862707138061523, 6.564248085021973, 4.892410755157471, 5.318929672241211, 6.568413257598877, 9.133810043334961, 7.237424373626709, 4.397247314453125, 7.976856708526611, 2.3416335582733154, 6.3428730964660645, 9.161764144897461, 9.220541000366211, 8.534843444824219, 4.4301581382751465, 3.7176058292388916, 4.751829624176025, 5.455737113952637, 0.894568920135498, 5.666994571685791, 1.3768253326416016, 9.591104507446289, 8.1528902053833, 7.843255519866943, 1.170852541923523, 9.404593467712402, 1.0466225147247314, 2.0946621894836426, 8.859201431274414, 7.388158321380615, 0.26116275787353516, 4.805743217468262, 7.201430320739746, 3.692185878753662, 4.12822151184082, 5.931211471557617, 6.904507637023926, 6.349696636199951, 4.302095413208008, 6.546483516693115, 8.069499969482422, 2.374738931655884, 9.065333366394043, 9.402444839477539, 6.593384265899658, 8.194018363952637, 3.910661458969116, 2.1357905864715576, 0.8497279286384583, 1.6840051412582397, 7.609548568725586, 1.4399187564849854, 1.0644301176071167, 8.14090347290039, 2.919234275817871, 2.675851583480835, 3.0641536712646484, 6.6956963539123535, 3.498112678527832, 0.69023597240448, 1.1889065504074097, 8.269831657409668, 8.937707901000977, 1.6059719324111938, 0.15695418417453766, 0.9913957118988037, 7.1435065269470215, 4.237936973571777, 2.295494556427002, 5.945472240447998, 3.4239118099212646, 4.425090312957764, 3.2222511768341064, 1.6377434730529785, 4.766551971435547, 6.3335795402526855, 3.467780351638794, 5.677145004272461, 1.9361212253570557, 2.3081438541412354, 5.466902256011963, 4.283600807189941, 2.0228371620178223, 5.984248638153076, 6.023360252380371, 8.454516410827637, 0.48479247093200684, 1.7577219009399414, 7.09525203704834, 1.0326943397521973},
	outsh:   []int{1, 5, 10, 16},
	outdata: []float64{3.588184118270874, 3.588184118270874, 6.765244483947754, 6.765244483947754, 8.703499794006348, 8.703499794006348, 7.832754611968994, 7.832754611968994, 7.205557346343994, 7.205557346343994, 2.2351908683776855, 2.2351908683776855, 9.140625, 9.140625, 1.8968262672424316, 1.8968262672424316, 3.588184118270874, 3.588184118270874, 6.765244483947754, 6.765244483947754, 8.703499794006348, 8.703499794006348, 7.832754611968994, 7.832754611968994, 7.205557346343994, 7.205557346343994, 2.2351908683776855, 2.2351908683776855, 9.140625, 9.140625, 1.8968262672424316, 1.8968262672424316, 7.579838752746582, 7.579838752746582, 2.0422589778900146, 2.0422589778900146, 1.7213886976242065, 1.7213886976242065, 8.008468627929688, 8.008468627929688, 1.8242313861846924, 1.8242313861846924, 0.15505588054656982, 0.15505588054656982, 7.218662738800049, 7.218662738800049, 6.959044456481934, 6.959044456481934, 7.579838752746582, 7.579838752746582, 2.0422589778900146, 2.0422589778900146, 1.7213886976242065, 1.7213886976242065, 8.008468627929688, 8.008468627929688, 1.8242313861846924, 1.8242313861846924, 0.15505588054656982, 0.15505588054656982, 7.218662738800049, 7.218662738800049, 6.959044456481934, 6.959044456481934, 5.268470287322998, 5.268470287322998, 3.0470921993255615, 3.0470921993255615, 6.904945373535156, 6.904945373535156, 0.048324376344680786, 0.048324376344680786, 5.862209796905518, 5.862209796905518, 9.752281188964844, 9.752281188964844, 9.729859352111816, 9.729859352111816, 6.917296409606934, 6.917296409606934, 5.268470287322998, 5.268470287322998, 3.0470921993255615, 3.0470921993255615, 6.904945373535156, 6.904945373535156, 0.048324376344680786, 0.048324376344680786, 5.862209796905518, 5.862209796905518, 9.752281188964844, 9.752281188964844, 9.729859352111816, 9.729859352111816, 6.917296409606934, 6.917296409606934, 5.983068466186523, 5.983068466186523, 8.5852689743042, 8.5852689743042, 8.877313613891602, 8.877313613891602, 0.06957779079675674, 0.06957779079675674, 8.360713958740234, 8.360713958740234, 1.0679903030395508, 1.0679903030395508, 1.2100963592529297, 1.2100963592529297, 9.838610649108887, 9.838610649108887, 5.983068466186523, 5.983068466186523, 8.5852689743042, 8.5852689743042, 8.877313613891602, 8.877313613891602, 0.06957779079675674, 0.06957779079675674, 8.360713958740234, 8.360713958740234, 1.0679903030395508, 1.0679903030395508, 1.2100963592529297, 1.2100963592529297, 9.838610649108887, 9.838610649108887, 2.814711809158325, 2.814711809158325, 6.38685941696167, 6.38685941696167, 8.260359764099121, 8.260359764099121, 2.532329559326172, 2.532329559326172, 7.589834213256836, 7.589834213256836, 1.294305443763733, 1.294305443763733, 0.6550663113594055, 0.6550663113594055, 0.8506288528442383, 0.8506288528442383, 2.814711809158325, 2.814711809158325, 6.38685941696167, 6.38685941696167, 8.260359764099121, 8.260359764099121, 2.532329559326172, 2.532329559326172, 7.589834213256836, 7.589834213256836, 1.294305443763733, 1.294305443763733, 0.6550663113594055, 0.6550663113594055, 0.8506288528442383, 0.8506288528442383, 2.1071043014526367, 2.1071043014526367, 8.425798416137695, 8.425798416137695, 8.408304214477539, 8.408304214477539, 7.394162178039551, 7.394162178039551, 2.590216875076294, 2.590216875076294, 0.7849066257476807, 0.7849066257476807, 2.676325559616089, 2.676325559616089, 8.072299003601074, 8.072299003601074, 2.1071043014526367, 2.1071043014526367, 8.425798416137695, 8.425798416137695, 8.408304214477539, 8.408304214477539, 7.394162178039551, 7.394162178039551, 2.590216875076294, 2.590216875076294, 0.7849066257476807, 0.7849066257476807, 2.676325559616089, 2.676325559616089, 8.072299003601074, 8.072299003601074, 3.4864327907562256, 3.4864327907562256, 2.422727346420288, 2.422727346420288, 9.429472923278809, 9.429472923278809, 2.553463935852051, 2.553463935852051, 5.96559476852417, 5.96559476852417, 7.482579231262207, 7.482579231262207, 2.3580377101898193, 2.3580377101898193, 1.5595018863677979, 1.5595018863677979, 3.4864327907562256, 3.4864327907562256, 2.422727346420288, 2.422727346420288, 9.429472923278809, 9.429472923278809, 2.553463935852051, 2.553463935852051, 5.96559476852417, 5.96559476852417, 7.482579231262207, 7.482579231262207, 2.3580377101898193, 2.3580377101898193, 1.5595018863677979, 1.5595018863677979, 0.034589122980833054, 0.034589122980833054, 3.0223019123077393, 3.0223019123077393, 4.7403154373168945, 4.7403154373168945, 0.2549790143966675, 0.2549790143966675, 6.765586853027344, 6.765586853027344, 6.91905403137207, 6.91905403137207, 1.7624300718307495, 1.7624300718307495, 1.8117458820343018, 1.8117458820343018, 0.034589122980833054, 0.034589122980833054, 3.0223019123077393, 3.0223019123077393, 4.7403154373168945, 4.7403154373168945, 0.2549790143966675, 0.2549790143966675, 6.765586853027344, 6.765586853027344, 6.91905403137207, 6.91905403137207, 1.7624300718307495, 1.7624300718307495, 1.8117458820343018, 1.8117458820343018, 5.739025592803955, 5.739025592803955, 7.67076301574707, 7.67076301574707, 8.772150039672852, 8.772150039672852, 2.1648266315460205, 2.1648266315460205, 2.8233001232147217, 2.8233001232147217, 3.455390214920044, 3.455390214920044, 1.954986333847046, 1.954986333847046, 6.2339935302734375, 6.2339935302734375, 5.739025592803955, 5.739025592803955, 7.67076301574707, 7.67076301574707, 8.772150039672852, 8.772150039672852, 2.1648266315460205, 2.1648266315460205, 2.8233001232147217, 2.8233001232147217, 3.455390214920044, 3.455390214920044, 1.954986333847046, 1.954986333847046, 6.2339935302734375, 6.2339935302734375, 8.180193901062012, 8.180193901062012, 3.4818568229675293, 3.4818568229675293, 0.5501848459243774, 0.5501848459243774, 7.886703014373779, 7.886703014373779, 3.614240884780884, 3.614240884780884, 7.023927211761475, 7.023927211761475, 0.4005505442619324, 0.4005505442619324, 2.0974090099334717, 2.0974090099334717, 8.180193901062012, 8.180193901062012, 3.4818568229675293, 3.4818568229675293, 0.5501848459243774, 0.5501848459243774, 7.886703014373779, 7.886703014373779, 3.614240884780884, 3.614240884780884, 7.023927211761475, 7.023927211761475, 0.4005505442619324, 0.4005505442619324, 2.0974090099334717, 2.0974090099334717, 6.166986465454102, 6.166986465454102, 4.413336753845215, 4.413336753845215, 7.7545647621154785, 7.7545647621154785, 5.656180381774902, 5.656180381774902, 8.241524696350098, 8.241524696350098, 0.1758895069360733, 0.1758895069360733, 7.527948379516602, 7.527948379516602, 2.4103832244873047, 2.4103832244873047, 6.166986465454102, 6.166986465454102, 4.413336753845215, 4.413336753845215, 7.7545647621154785, 7.7545647621154785, 5.656180381774902, 5.656180381774902, 8.241524696350098, 8.241524696350098, 0.1758895069360733, 0.1758895069360733, 7.527948379516602, 7.527948379516602, 2.4103832244873047, 2.4103832244873047, 2.5865209102630615, 2.5865209102630615, 8.264690399169922, 8.264690399169922, 4.0821614265441895, 4.0821614265441895, 4.711780548095703, 4.711780548095703, 8.985666275024414, 8.985666275024414, 3.614298105239868, 3.614298105239868, 2.9979336261749268, 2.9979336261749268, 1.0938977003097534, 1.0938977003097534, 2.5865209102630615, 2.5865209102630615, 8.264690399169922, 8.264690399169922, 4.0821614265441895, 4.0821614265441895, 4.711780548095703, 4.711780548095703, 8.985666275024414, 8.985666275024414, 3.614298105239868, 3.614298105239868, 2.9979336261749268, 2.9979336261749268, 1.0938977003097534, 1.0938977003097534, 0.6748584508895874, 0.6748584508895874, 8.776729583740234, 8.776729583740234, 7.792627811431885, 7.792627811431885, 4.120683670043945, 4.120683670043945, 8.067517280578613, 8.067517280578613, 6.481339931488037, 6.481339931488037, 3.7537691593170166, 3.7537691593170166, 1.5674660205841064, 1.5674660205841064, 0.6748584508895874, 0.6748584508895874, 8.776729583740234, 8.776729583740234, 7.792627811431885, 7.792627811431885, 4.120683670043945, 4.120683670043945, 8.067517280578613, 8.067517280578613, 6.481339931488037, 6.481339931488037, 3.7537691593170166, 3.7537691593170166, 1.5674660205841064, 1.5674660205841064, 4.622176647186279, 4.622176647186279, 6.572947025299072, 6.572947025299072, 6.97075891494751, 6.97075891494751, 0.5755885243415833, 0.5755885243415833, 9.231983184814453, 9.231983184814453, 5.7552490234375, 5.7552490234375, 6.862707138061523, 6.862707138061523, 6.564248085021973, 6.564248085021973, 4.622176647186279, 4.622176647186279, 6.572947025299072, 6.572947025299072, 6.97075891494751, 6.97075891494751, 0.5755885243415833, 0.5755885243415833, 9.231983184814453, 9.231983184814453, 5.7552490234375, 5.7552490234375, 6.862707138061523, 6.862707138061523, 6.564248085021973, 6.564248085021973, 4.892410755157471, 4.892410755157471, 5.318929672241211, 5.318929672241211, 6.568413257598877, 6.568413257598877, 9.133810043334961, 9.133810043334961, 7.237424373626709, 7.237424373626709, 4.397247314453125, 4.397247314453125, 7.976856708526611, 7.976856708526611, 2.3416335582733154, 2.3416335582733154, 4.892410755157471, 4.892410755157471, 5.318929672241211, 5.318929672241211, 6.568413257598877, 6.568413257598877, 9.133810043334961, 9.133810043334961, 7.237424373626709, 7.237424373626709, 4.397247314453125, 4.397247314453125, 7.976856708526611, 7.976856708526611, 2.3416335582733154, 2.3416335582733154, 6.3428730964660645, 6.3428730964660645, 9.161764144897461, 9.161764144897461, 9.220541000366211, 9.220541000366211, 8.534843444824219, 8.534843444824219, 4.4301581382751465, 4.4301581382751465, 3.7176058292388916, 3.7176058292388916, 4.751829624176025, 4.751829624176025, 5.455737113952637, 5.455737113952637, 6.3428730964660645, 6.3428730964660645, 9.161764144897461, 9.161764144897461, 9.220541000366211, 9.220541000366211, 8.534843444824219, 8.534843444824219, 4.4301581382751465, 4.4301581382751465, 3.7176058292388916, 3.7176058292388916, 4.751829624176025, 4.751829624176025, 5.455737113952637, 5.455737113952637, 0.894568920135498, 0.894568920135498, 5.666994571685791, 5.666994571685791, 1.3768253326416016, 1.3768253326416016, 9.591104507446289, 9.591104507446289, 8.1528902053833, 8.1528902053833, 7.843255519866943, 7.843255519866943, 1.170852541923523, 1.170852541923523, 9.404593467712402, 9.404593467712402, 0.894568920135498, 0.894568920135498, 5.666994571685791, 5.666994571685791, 1.3768253326416016, 1.3768253326416016, 9.591104507446289, 9.591104507446289, 8.1528902053833, 8.1528902053833, 7.843255519866943, 7.843255519866943, 1.170852541923523, 1.170852541923523, 9.404593467712402, 9.404593467712402, 1.0466225147247314, 1.0466225147247314, 2.0946621894836426, 2.0946621894836426, 8.859201431274414, 8.859201431274414, 7.388158321380615, 7.388158321380615, 0.26116275787353516, 0.26116275787353516, 4.805743217468262, 4.805743217468262, 7.201430320739746, 7.201430320739746, 3.692185878753662, 3.692185878753662, 1.0466225147247314, 1.0466225147247314, 2.0946621894836426, 2.0946621894836426, 8.859201431274414, 8.859201431274414, 7.388158321380615, 7.388158321380615, 0.26116275787353516, 0.26116275787353516, 4.805743217468262, 4.805743217468262, 7.201430320739746, 7.201430320739746, 3.692185878753662, 3.692185878753662, 4.12822151184082, 4.12822151184082, 5.931211471557617, 5.931211471557617, 6.904507637023926, 6.904507637023926, 6.349696636199951, 6.349696636199951, 4.302095413208008, 4.302095413208008, 6.546483516693115, 6.546483516693115, 8.069499969482422, 8.069499969482422, 2.374738931655884, 2.374738931655884, 4.12822151184082, 4.12822151184082, 5.931211471557617, 5.931211471557617, 6.904507637023926, 6.904507637023926, 6.349696636199951, 6.349696636199951, 4.302095413208008, 4.302095413208008, 6.546483516693115, 6.546483516693115, 8.069499969482422, 8.069499969482422, 2.374738931655884, 2.374738931655884, 9.065333366394043, 9.065333366394043, 9.402444839477539, 9.402444839477539, 6.593384265899658, 6.593384265899658, 8.194018363952637, 8.194018363952637, 3.910661458969116, 3.910661458969116, 2.1357905864715576, 2.1357905864715576, 0.8497279286384583, 0.8497279286384583, 1.6840051412582397, 1.6840051412582397, 9.065333366394043, 9.065333366394043, 9.402444839477539, 9.402444839477539, 6.593384265899658, 6.593384265899658, 8.194018363952637, 8.194018363952637, 3.910661458969116, 3.910661458969116, 2.1357905864715576, 2.1357905864715576, 0.8497279286384583, 0.8497279286384583, 1.6840051412582397, 1.6840051412582397, 7.609548568725586, 7.609548568725586, 1.4399187564849854, 1.4399187564849854, 1.0644301176071167, 1.0644301176071167, 8.14090347290039, 8.14090347290039, 2.919234275817871, 2.919234275817871, 2.675851583480835, 2.675851583480835, 3.0641536712646484, 3.0641536712646484, 6.6956963539123535, 6.6956963539123535, 7.609548568725586, 7.609548568725586, 1.4399187564849854, 1.4399187564849854, 1.0644301176071167, 1.0644301176071167, 8.14090347290039, 8.14090347290039, 2.919234275817871, 2.919234275817871, 2.675851583480835, 2.675851583480835, 3.0641536712646484, 3.0641536712646484, 6.6956963539123535, 6.6956963539123535, 3.498112678527832, 3.498112678527832, 0.69023597240448, 0.69023597240448, 1.1889065504074097, 1.1889065504074097, 8.269831657409668, 8.269831657409668, 8.937707901000977, 8.937707901000977, 1.6059719324111938, 1.6059719324111938, 0.15695418417453766, 0.15695418417453766, 0.9913957118988037, 0.9913957118988037, 3.498112678527832, 3.498112678527832, 0.69023597240448, 0.69023597240448, 1.1889065504074097, 1.1889065504074097, 8.269831657409668, 8.269831657409668, 8.937707901000977, 8.937707901000977, 1.6059719324111938, 1.6059719324111938, 0.15695418417453766, 0.15695418417453766, 0.9913957118988037, 0.9913957118988037, 7.1435065269470215, 7.1435065269470215, 4.237936973571777, 4.237936973571777, 2.295494556427002, 2.295494556427002, 5.945472240447998, 5.945472240447998, 3.4239118099212646, 3.4239118099212646, 4.425090312957764, 4.425090312957764, 3.2222511768341064, 3.2222511768341064, 1.6377434730529785, 1.6377434730529785, 7.1435065269470215, 7.1435065269470215, 4.237936973571777, 4.237936973571777, 2.295494556427002, 2.295494556427002, 5.945472240447998, 5.945472240447998, 3.4239118099212646, 3.4239118099212646, 4.425090312957764, 4.425090312957764, 3.2222511768341064, 3.2222511768341064, 1.6377434730529785, 1.6377434730529785, 4.766551971435547, 4.766551971435547, 6.3335795402526855, 6.3335795402526855, 3.467780351638794, 3.467780351638794, 5.677145004272461, 5.677145004272461, 1.9361212253570557, 1.9361212253570557, 2.3081438541412354, 2.3081438541412354, 5.466902256011963, 5.466902256011963, 4.283600807189941, 4.283600807189941, 4.766551971435547, 4.766551971435547, 6.3335795402526855, 6.3335795402526855, 3.467780351638794, 3.467780351638794, 5.677145004272461, 5.677145004272461, 1.9361212253570557, 1.9361212253570557, 2.3081438541412354, 2.3081438541412354, 5.466902256011963, 5.466902256011963, 4.283600807189941, 4.283600807189941, 2.0228371620178223, 2.0228371620178223, 5.984248638153076, 5.984248638153076, 6.023360252380371, 6.023360252380371, 8.454516410827637, 8.454516410827637, 0.48479247093200684, 0.48479247093200684, 1.7577219009399414, 1.7577219009399414, 7.09525203704834, 7.09525203704834, 1.0326943397521973, 1.0326943397521973, 2.0228371620178223, 2.0228371620178223, 5.984248638153076, 5.984248638153076, 6.023360252380371, 6.023360252380371, 8.454516410827637, 8.454516410827637, 0.48479247093200684, 0.48479247093200684, 1.7577219009399414, 1.7577219009399414, 7.09525203704834, 7.09525203704834, 1.0326943397521973, 1.0326943397521973},
}

func TestUpsampleSimple(t *testing.T) {
	tt := tensor.New(tensor.Of(tensor.Float64), tensor.WithShape(1, 1, 3, 3), tensor.WithBacking([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	upop := newUpsampleOp(tt.Shape(), 2)
	fmt.Println(tt)
	out, err := upop.Do(tt)
	t.Log(out)
	if err != nil {
		panic(err)
	}
}

func TestUpsampleWithData(t *testing.T) {
	tt := tensor.New(tensor.Of(tensor.Float64), tensor.WithShape(data.insh...), tensor.WithBacking(data.indata))
	expected := tensor.New(tensor.Of(tensor.Float64), tensor.WithShape(data.outsh...), tensor.WithBacking(data.outdata))

	g := NewGraph()
	inp := NewTensor(g, tensor.Float64, 4, WithShape(data.insh...), WithName("inp"))
	out := Must(Upsample2D(inp, 2))

	vm := NewTapeMachine(g)
	if err := Let(inp, tt); err != nil {
		panic(err)
	}
	vm.RunAll()
	// fmt.Println(out.Value())
	vm.Close()
	assert.Equal(t, expected.Data(), out.Value().(*tensor.Dense).Data(), "Output is not equal to expected value")

}